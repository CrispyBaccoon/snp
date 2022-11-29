package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// FilterValue is the snippet filter value that can be used when searching.
func (s Snippet) FilterValue() string {
	return s.Folder + "/" + s.Name + "." + s.Language
}

// snippetDelegate represents the snippet list item.
type snippetDelegate struct {
	styles SnippetsBaseStyle
	state  state
}

// Height is the number of lines the snippet list item takes up.
func (d snippetDelegate) Height() int {
	return 2
}

// Spacing is the number of lines to insert between list items.
func (d snippetDelegate) Spacing() int {
	return 1
}

// Update is called when the list is updated.
// We use this to update the snippet code view.
func (d snippetDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return func() tea.Msg {
		if m.SelectedItem() == nil {
			return nil
		}
		return updateContentMsg(m.SelectedItem().(Snippet))
	}
}

// Render renders the list item for the snippet which includes the title,
// folder, and date.
func (d snippetDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	if item == nil {
		return
	}
	s, ok := item.(Snippet)
	if !ok {
		return
	}

	titleStyle := d.styles.SelectedTitle
	subtitleStyle := d.styles.SelectedSubtitle
	if d.state == copyingState {
		titleStyle = d.styles.CopiedTitle
		subtitleStyle = d.styles.CopiedSubtitle
	} else if d.state == deletingState {
		titleStyle = d.styles.DeletedTitle
		subtitleStyle = d.styles.DeletedSubtitle
	}

	if index == m.Index() {
		fmt.Fprintln(w, "  "+titleStyle.Render(s.Name))
		fmt.Fprint(w, "  "+subtitleStyle.Render(s.Folder + " • " + s.Language))
		return
	}
	fmt.Fprintln(w, "  "+d.styles.UnselectedTitle.Render(s.Name))
	fmt.Fprint(w, "  "+d.styles.UnselectedSubtitle.Render(s.Folder + " • " + s.Language))
}

// Folder represents a group of snippets in a directory.
type Folder string

// FilterValue is the searchable value for the folder.
func (f Folder) FilterValue() string {
	return string(f)
}

// folderDelegate represents a folder list item.
type folderDelegate struct{ styles FoldersBaseStyle }

// Height is the number of lines the folder list item takes up.
func (d folderDelegate) Height() int {
	return 1
}

// Spacing is the number of lines to insert between folder items.
func (d folderDelegate) Spacing() int {
	return 0
}

// Update is what is called when the folder selection is updated.
// TODO: Update the filter search for the snippets with the folder name.
func (d folderDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

// Render renders a folder list item.
func (d folderDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	f, ok := item.(Folder)
	if !ok {
		return
	}
	fmt.Fprint(w, "  ")
	if index == m.Index() {
		fmt.Fprint(w, d.styles.Selected.Render("• " + string(f)))
		return
	}
	fmt.Fprint(w, d.styles.Unselected.Render("  " + string(f)))
}

