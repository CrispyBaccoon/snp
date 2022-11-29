package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alecthomas/chroma/v2/quick"
)

// TODO:
// default values for empty state.
const defaultSnippetFolder = "misc"
const defaultSnippetName = "Untitled"
const defaultSnippetFileName = "snippet.txt"
const defaultLanguage = "go"

// defaultSnippet is a snippet with all of the default values, used for when
// there are no snippets available.
var defaultSnippet = Snippet{
	Folder:   defaultSnippetFolder,
	Name:     defaultSnippetName,
	File:     defaultSnippetFileName,
	Language: defaultLanguage,
}

// Snippet represents a snippet of code in a language.
// It is nested within a folder
type Snippet struct {
	Folder   string
	Name     string
	File     string
	Language string
}

// String returns the folder/name.ext of the snippet.
func (s Snippet) String() string {
	return fmt.Sprintf("%s/%s.%s", s.Folder, s.Name, s.Language)
}

// Content returns the snippet contents.
func (s Snippet) Content(highlight bool) string {
	config := readConfig()
	file := filepath.Join(config.Root, s.Folder, s.File)
	content, err := os.ReadFile(file)
	if err != nil {
		return ""
	}

	if !highlight {
		return string(content)
	}

	var b bytes.Buffer
	err = quick.Highlight(&b, string(content), s.Language, "terminal16m", config.Theme)
	if err != nil {
		return string(content)
	}
	return b.String()
}
