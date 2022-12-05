package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
	"github.com/caarlos0/env/v6"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
	"github.com/sahilm/fuzzy"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

func main() {
	config := readConfig()
	snippets := readSnippets(config)
	stdin := readStdin()
	if stdin != "" {
		saveSnippet(stdin, config, snippets)
		return
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "list":
			listSnippets(snippets)
		default:
			snippet := findSnippet(os.Args[1], snippets)
			fmt.Print(snippet.Content(isatty.IsTerminal(os.Stdout.Fd())))
		}
		return
	}

	err := runInteractiveMode(config, snippets)
	if err != nil {
		fmt.Println("Alas, there's been an error", err)
	}
}

// parseName returns a folder, name, and language for the given name.
// this is useful for parsing file names when passed as command line arguments.
//
// Example:
//
//	Notes/Hello.go -> (Notes, Hello, go)
//	Hello.go       -> (Misc, Hello, go)
//	Notes/Hello    -> (Notes, Hello, go)
func parseName(s string) (string, string, string) {
	var (
		folder    = defaultSnippetFolder
		name      = defaultSnippetName
		language  = defaultLanguage
		remaining string
	)

	tokens := strings.Split(s, "/")
	if len(tokens) > 1 {
		folder = tokens[0]
		remaining = tokens[1]
	} else {
		remaining = tokens[0]
	}

	tokens = strings.Split(remaining, ".")
	if len(tokens) > 1 {
		name = tokens[0]
		language = tokens[1]
	} else {
		name = tokens[0]
	}

	return folder, name, language
}

// readStdin returns the stdin that is piped in to the command line interface.
func readStdin() string {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return ""
	}

	if stat.Mode()&os.ModeCharDevice != 0 {
		return ""
	}

	reader := bufio.NewReader(os.Stdin)
	var b strings.Builder

	for {
		r, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		_, err = b.WriteRune(r)
		if err != nil {
			return ""
		}
	}

	return b.String()
}

// defaultConfig returns the default config path
func defaultConfig() string {
	if c := os.Getenv("SNP_CONFIG"); c != "" {
		return c
	}
	cfgPath, err := xdg.ConfigFile("snp/config.yaml")
	if err != nil {
		return "config.yaml"
	}
	return cfgPath
}

// readConfig returns a configuration read from the environment.
func readConfig() Config {
	config := newConfig()
	fi, err := os.Open(defaultConfig())
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return newConfig()
	}
	if fi != nil {
		defer fi.Close()
		if err := yaml.NewDecoder(fi).Decode(&config); err != nil {
			return newConfig()
		}
	}

	if err := env.Parse(&config); err != nil {
		return newConfig()
	}

	return config
}

// TODO:
// readSnippets returns all the snippets read from the snippets.json file.
func readSnippets(config Config) []Snippet {
	var snippets []Snippet
	fd, err := ioutil.ReadDir(config.Root)
	if err != nil {
		return snippets
	}

	parseFile := func(d fs.FileInfo, p string) {
		if d.IsDir() {
			return
		}
		fname := d.Name()
		str := strings.Split(fname, ".")
		var name string
		lan := "txt"
		if len(str) < 2 {
			name = str[0]
		} else {
			name = strings.Join(str[:len(str)-1], ".")
			lan = str[len(str)-1]
		}
		snippets = append(snippets, Snippet{Name: name, Folder: p, Language: lan})
	}

	parseDir := func(d fs.FileInfo) {
		if !d.IsDir() {
			return
		}
		fdd, err := ioutil.ReadDir(filepath.Join(config.Root, d.Name()))
		if err != nil {
			return
		}
		for _, dd := range fdd {
			if !dd.IsDir() {
				parseFile(dd, d.Name())
			}
		}
	}

	for _, d := range fd {
		if !d.IsDir() {
			parseFile(d, defaultSnippetFolder)
		} else {
			parseDir(d)
		}
	}
	return snippets
}

func saveSnippet(content string, config Config, snippets []Snippet) {
	// Save snippet to location
	var name string = defaultSnippetName
	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	}

	folder, name, language := parseName(name)
	file := fmt.Sprintf("%s.%s", name, language)
	err := os.WriteFile(filepath.Join(config.Root, folder, file), []byte(content), 0644)
	if err != nil {
		fmt.Println("unable to create snippet")
		return
	}
}

func listSnippets(snippets []Snippet) {
	for _, snippet := range snippets {
		fmt.Println(snippet)
	}
}

// Snippets is a wrapper for a snippets array to implement the fuzzy.Source
// interface.
type Snippets struct {
	snippets []Snippet
}

// String returns the string of the snippet at the specified position i
func (s Snippets) String(i int) string {
	return s.snippets[i].String()
}

// Len returns the length of the snippets array.
func (s Snippets) Len() int {
	return len(s.snippets)
}

func findSnippet(search string, snippets []Snippet) Snippet {
	matches := fuzzy.FindFrom(os.Args[1], Snippets{snippets})
	if len(matches) > 0 {
		return snippets[matches[0].Index]
	}
	return Snippet{}
}

func runInteractiveMode(config Config, snippets []Snippet) error {
	var folders = make(map[Folder][]list.Item)
	var items []list.Item
	for _, snippet := range snippets {
		folders[Folder(snippet.Folder)] = append(folders[Folder(snippet.Folder)], list.Item(snippet))
	}
	if len(items) <= 0 {
		items = append(items, list.Item(defaultSnippet))
	}

	defaultStyles := DefaultStyles(config)

	var folderItems []list.Item
	foldersSlice := maps.Keys(folders)
	slices.Sort(foldersSlice)
	for _, folder := range foldersSlice {
		folderItems = append(folderItems, list.Item(Folder(folder)))
	}
	if len(folderItems) <= 0 {
		folderItems = append(folderItems, list.Item(Folder(defaultSnippetFolder)))
	}
	folderList := list.New(folderItems, folderDelegate{defaultStyles.Folders.Blurred}, 0, 0)
	folderList.Title = "Folders"

	folderList.SetShowHelp(false)
	folderList.SetFilteringEnabled(false)
	folderList.SetShowStatusBar(false)
	folderList.DisableQuitKeybindings()
	folderList.Styles.NoItems = lipgloss.NewStyle().Margin(0, 2).Foreground(lipgloss.Color(config.GrayColor))
	folderList.SetStatusBarItemName("folder", "folders")

	content := viewport.New(80, 0)

	lists := map[Folder]*list.Model{}

	for folder, items := range folders {
		lists[folder] = newList(items, 20, defaultStyles.Snippets.Focused)
	}

	m := &Model{
		Lists:        lists,
		Folders:      folderList,
		Code:         content,
		ContentStyle: defaultStyles.Content.Blurred,
		ListStyle:    defaultStyles.Snippets.Focused,
		FoldersStyle: defaultStyles.Folders.Blurred,
		keys:         DefaultKeyMap,
		help:         help.New(),
		config:       config,
		inputs: []textinput.Model{
			newTextInput(defaultSnippetFolder + " "),
			newTextInput(defaultSnippetName + " "),
			newTextInput(config.DefaultLanguage),
		},
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	model, err := p.Run()
	if err != nil {
		return err
	}
	fm, ok := model.(*Model)
	if !ok {
		return err
	}
	var allSnippets []list.Item
	for _, list := range fm.Lists {
		allSnippets = append(allSnippets, list.Items()...)
	}
	if len(allSnippets) <= 0 {
		allSnippets = []list.Item{defaultSnippet}
	}
	/* b, err := json.Marshal(allSnippets)
	if err != nil {
		return err
	} */
	// err = os.WriteFile(filepath.Join(config.Home, config.File), b, os.ModePerm)
	/* if err != nil {
		return err
	} */
	return nil
}

func newList(items []list.Item, height int, styles SnippetsBaseStyle) *list.Model {
	snippetList := list.New(items, snippetDelegate{styles, navigatingState}, 25, height)
	snippetList.SetShowHelp(false)
	snippetList.SetShowFilter(false)
	snippetList.SetShowTitle(false)
	snippetList.Styles.StatusBar = lipgloss.NewStyle().Margin(1, 2).Foreground(lipgloss.Color("240")).MaxWidth(35 - 2)
	snippetList.Styles.NoItems = lipgloss.NewStyle().Margin(0, 2).Foreground(lipgloss.Color("8")).MaxWidth(35 - 2)
	snippetList.FilterInput.Prompt = "Find: "
	snippetList.FilterInput.PromptStyle = styles.Title
	snippetList.SetStatusBarItemName("snippet", "snippets")
	snippetList.DisableQuitKeybindings()
	snippetList.Styles.Title = styles.Title
	snippetList.Styles.TitleBar = styles.TitleBar

	return &snippetList
}

func newTextInput(placeholder string) textinput.Model {
	i := textinput.New()
	i.Prompt = ""
	i.PromptStyle = lipgloss.NewStyle().Margin(0, 1)
	i.Placeholder = placeholder
	return i
}
