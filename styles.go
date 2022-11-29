package main

import "github.com/charmbracelet/lipgloss"

// SnippetsStyle is the style struct to handle the focusing and blurring of the
// snippets pane in the application.
type SnippetsStyle struct {
	Focused SnippetsBaseStyle
	Blurred SnippetsBaseStyle
}

// FoldersStyle is the style struct to handle the focusing and blurring of the
// folders pane in the application.
type FoldersStyle struct {
	Focused FoldersBaseStyle
	Blurred FoldersBaseStyle
}

// ContentStyle is the style struct to handle the focusing and blurring of the
// content pane in the application.
type ContentStyle struct {
	Focused ContentBaseStyle
	Blurred ContentBaseStyle
}

// SnippetsBaseStyle holds the neccessary styling for the snippets pane of
// the application.
type SnippetsBaseStyle struct {
	Base               lipgloss.Style
	Title              lipgloss.Style
	TitleBar           lipgloss.Style
	SelectedSubtitle   lipgloss.Style
	UnselectedSubtitle lipgloss.Style
	SelectedTitle      lipgloss.Style
	UnselectedTitle    lipgloss.Style
	CopiedTitleBar     lipgloss.Style
	CopiedTitle        lipgloss.Style
	CopiedSubtitle     lipgloss.Style
	DeletedTitleBar    lipgloss.Style
	DeletedTitle       lipgloss.Style
	DeletedSubtitle    lipgloss.Style
}

// FoldersBaseStyle holds the neccessary styling for the folders pane of
// the application.
type FoldersBaseStyle struct {
	Base       lipgloss.Style
	Title      lipgloss.Style
	TitleBar   lipgloss.Style
	Selected   lipgloss.Style
	Unselected lipgloss.Style
}

// ContentBaseStyle holds the neccessary styling for the content pane of the
// application.
type ContentBaseStyle struct {
	Base         lipgloss.Style
	Title        lipgloss.Style
	Separator    lipgloss.Style
	LineNumber   lipgloss.Style
	EmptyHint    lipgloss.Style
	EmptyHintKey lipgloss.Style
}

// Styles is the struct of all styles for the application.
type Styles struct {
	Snippets SnippetsStyle
	Folders  FoldersStyle
	Content  ContentStyle
}

var marginStyle = lipgloss.NewStyle().Margin(1, 0, 0, 1)

// DefaultStyles is the default implementation of the styles struct for all
// styling in the application.
func DefaultStyles(config Config) Styles {
	white := lipgloss.Color(config.ForegroundColor)
	black := lipgloss.Color(config.BackgroundColor)
	red := lipgloss.Color(config.RedColor)
	green := lipgloss.Color(config.GreenColor)
	// yellow := lipgloss.Color(config.YellowColor)
	blue := lipgloss.Color(config.BlueColor)
	// magenta := lipgloss.Color(config.MagentaColor)
	// cyan := lipgloss.Color(config.CyanColor)
	brightRed := lipgloss.Color(config.BrightRedColor)
	brightGreen := lipgloss.Color(config.BrightGreenColor)
	// brightYellow := lipgloss.Color(config.BrightYellowColor)
	brightBlue := lipgloss.Color(config.BrightBlueColor)
	// brightMagenta := lipgloss.Color(config.BrightMagentaColor)
	// brightCyan := lipgloss.Color(config.BrightCyanColor)
	gray := lipgloss.Color(config.GrayColor)

	return Styles{
		Snippets: SnippetsStyle{
			Focused: SnippetsBaseStyle{
				Base:               lipgloss.NewStyle().Width(35),
				TitleBar:           lipgloss.NewStyle().Background(blue).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1).Foreground(white),
				SelectedSubtitle:   lipgloss.NewStyle().Foreground(blue),
				UnselectedSubtitle: lipgloss.NewStyle().Foreground(lipgloss.Color("237")),
				SelectedTitle:      lipgloss.NewStyle().Foreground(brightBlue),
				UnselectedTitle:    lipgloss.NewStyle().Foreground(gray),
				CopiedTitleBar:     lipgloss.NewStyle().Background(green).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1).Foreground(white),
				CopiedTitle:        lipgloss.NewStyle().Foreground(brightGreen),
				CopiedSubtitle:     lipgloss.NewStyle().Foreground(green),
				DeletedTitleBar:    lipgloss.NewStyle().Background(red).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1).Foreground(white),
				DeletedTitle:       lipgloss.NewStyle().Foreground(brightRed),
				DeletedSubtitle:    lipgloss.NewStyle().Foreground(red),
			},
			Blurred: SnippetsBaseStyle{
				Base:               lipgloss.NewStyle().Width(35),
				TitleBar:           lipgloss.NewStyle().Background(black).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1).Foreground(gray),
				SelectedSubtitle:   lipgloss.NewStyle().Foreground(blue),
				UnselectedSubtitle: lipgloss.NewStyle().Foreground(black),
				SelectedTitle:      lipgloss.NewStyle().Foreground(brightBlue),
				UnselectedTitle:    lipgloss.NewStyle().Foreground(lipgloss.Color("237")),
				CopiedTitleBar:     lipgloss.NewStyle().Background(green).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1),
				CopiedTitle:        lipgloss.NewStyle().Foreground(brightGreen),
				CopiedSubtitle:     lipgloss.NewStyle().Foreground(green),
				DeletedTitleBar:    lipgloss.NewStyle().Background(red).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1),
				DeletedTitle:       lipgloss.NewStyle().Foreground(brightRed),
				DeletedSubtitle:    lipgloss.NewStyle().Foreground(red),
			},
		},
		Folders: FoldersStyle{
			Focused: FoldersBaseStyle{
				Base:       lipgloss.NewStyle().Width(22),
				Title:      lipgloss.NewStyle().Padding(0, 1).Foreground(white),
				TitleBar:   lipgloss.NewStyle().Background(blue).Width(22-2).Margin(0, 1, 1, 1),
				Selected:   lipgloss.NewStyle().Foreground(brightBlue),
				Unselected: lipgloss.NewStyle().Foreground(gray),
			},
			Blurred: FoldersBaseStyle{
				Base:       lipgloss.NewStyle().Width(22),
				Title:      lipgloss.NewStyle().Padding(0, 1).Foreground(gray),
				TitleBar:   lipgloss.NewStyle().Background(black).Width(22-2).Margin(0, 1, 1, 1),
				Selected:   lipgloss.NewStyle().Foreground(brightBlue),
				Unselected: lipgloss.NewStyle().Foreground(lipgloss.Color("237")),
			},
		},
		Content: ContentStyle{
			Focused: ContentBaseStyle{
				Base:         lipgloss.NewStyle().Margin(0, 1),
				Title:        lipgloss.NewStyle().Background(blue).Foreground(white).Margin(0, 0, 1, 1).Padding(0, 1),
				Separator:    lipgloss.NewStyle().Foreground(white).Margin(0, 0, 1, 1),
				LineNumber:   lipgloss.NewStyle().Foreground(gray),
				EmptyHint:    lipgloss.NewStyle().Foreground(gray),
				EmptyHintKey: lipgloss.NewStyle().Foreground(brightBlue),
			},
			Blurred: ContentBaseStyle{
				Base:         lipgloss.NewStyle().Margin(0, 1),
				Title:        lipgloss.NewStyle().Background(black).Foreground(gray).Margin(0, 0, 1, 1).Padding(0, 1),
				Separator:    lipgloss.NewStyle().Foreground(gray).Margin(0, 0, 1, 1),
				LineNumber:   lipgloss.NewStyle().Foreground(black),
				EmptyHint:    lipgloss.NewStyle().Foreground(gray),
				EmptyHintKey: lipgloss.NewStyle().Foreground(brightBlue),
			},
		},
	}
}
