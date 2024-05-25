// ui/home.go

package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define styles
var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

type operation struct {
	title, desc string
}

func (i operation) Title() string       { return i.title }
func (i operation) Description() string { return i.desc }
func (i operation) FilterValue() string { return i.title }

type HomeUI struct {
	options list.Model
	login   loginForm // Add login form as a field
	mode    string    // Add a mode field to track current UI state
}

func NewHomeUI() HomeUI {
	options := []list.Item{
		operation{title: "Login", desc: "Login to your existing pwdManager account"},
		operation{title: "Signup", desc: "Create a new account to start using pwdManager"},
	}

	ui := HomeUI{
		options: list.New(options, list.NewDefaultDelegate(), 0, 0),
		login:   initialLoginForm(),
		mode:    "home", // Initialize in home mode
	}

	ui.options.Title = "Welcome to pwdManager!\nPlease choose an option to start using:"
	ui.options.SetShowHelp(true) // Ensure help text is shown

	return ui
}

func (m HomeUI) Init() tea.Cmd {
	return nil
}

func (m HomeUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.mode {
	case "home":
		return m.updateHome(msg)
	case "login":
		return m.updateLogin(msg)
	default:
		return m, nil
	}
}

func (m HomeUI) updateHome(msg tea.Msg) (HomeUI, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			i, ok := m.options.SelectedItem().(operation)
			if ok && i.title == "Login" {
				m.mode = "login" // Switch to login mode
				return m, m.login.Init()
			}
			if ok && i.title == "Signup" {
				// Handle signup logic
				return m, nil
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.options.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.options, cmd = m.options.Update(msg)
	return m, cmd
}

func (m HomeUI) updateLogin(msg tea.Msg) (HomeUI, tea.Cmd) {
	if strMsg, ok := msg.(string); ok && strMsg == "back-to-home" {
		m.mode = "home"
		return m, nil
	}

	model, cmd := m.login.Update(msg)
	m.login = model.(loginForm) // Type assertion
	return m, cmd
}

func (m HomeUI) View() string {
	switch m.mode {
	case "home":
		return docStyle.Render(m.options.View())
	case "login":
		return m.login.View()
	default:
		return ""
	}
}

// Ensure HomeUI implements the model interface
var _ model = (*HomeUI)(nil)
