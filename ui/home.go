package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type operation struct {
	title, desc string
}

func (i operation) Title() string       { return i.title }
func (i operation) Description() string { return i.desc }
func (i operation) FilterValue() string { return i.title }

type HomeUI struct {
	options list.Model
}

func (m HomeUI) Init() tea.Cmd {
	return nil
}

func (m HomeUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.options.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.options, cmd = m.options.Update(msg)
	return m, cmd
}

func (m HomeUI) View() string {
	return docStyle.Render(m.options.View())
}

func NewHomeUI() *HomeUI {
	options := []list.Item{
		operation{title: "Login", desc: "Login to your existing pwdManager account"},
		operation{title: "Signup", desc: "Create a new account to start using pwdManager"},
	}

	ui := &HomeUI{
		options: list.New(options, list.NewDefaultDelegate(), 0, 0),
	}

	ui.options.Title = "Welcome to pwdManager!\nPlease choose an option to start using:"

	return ui
}
