// ui/login.go

package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	noStyle          = lipgloss.NewStyle()
	placeholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	helpStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type loginForm struct {
	emailInput    textinput.Model
	passwordInput textinput.Model
	err           error
}

func initialLoginForm() loginForm {
	email := textinput.New()
	email.Prompt = "Email:"
	email.Placeholder = "Email"
	email.Focus()
	email.CharLimit = 80
	email.Width = 30

	email.PromptStyle = focusStyle
	email.TextStyle = noStyle
	email.PlaceholderStyle = placeholderStyle
	email.Cursor.Style = cursorStyle

	password := textinput.New()
	password.Prompt = "Password:"
	password.Placeholder = "Password"
	password.EchoMode = textinput.EchoPassword
	password.CharLimit = 80
	password.EchoCharacter = '•'
	password.Width = 30

	password.PromptStyle = focusStyle
	password.TextStyle = noStyle
	password.PlaceholderStyle = placeholderStyle
	password.Cursor.Style = cursorStyle

	return loginForm{
		emailInput:    email,
		passwordInput: password,
	}
}

func (l loginForm) Init() tea.Cmd {
	return textinput.Blink
}

func (l loginForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Perform login
			print("pressed enter")
		case "tab":
			if l.emailInput.Focused() {
				l.emailInput.Blur()
				l.passwordInput.Focus()
			} else {
				l.passwordInput.Blur()
				l.emailInput.Focus()
			}
		case "esc":
			return l, func() tea.Msg {
				return tea.Msg("back-to-home")
			}
		}
	}

	_, emailCmd := l.emailInput.Update(msg)
	_, passwordCmd := l.passwordInput.Update(msg)

	cmds = append(cmds, emailCmd, passwordCmd)

	return l, tea.Batch(cmds...)
}

func (l loginForm) View() string {
	hint := helpView()
	if l.err != nil {
		return docStyle.Render(fmt.Sprintf(
			"Email:\n%s\n\nPassword:\n%s\n\nError: %v\n\n(Press Enter to Login)%s",
			l.emailInput.View(),
			l.passwordInput.View(),
			l.err,
			hint,
		))
	}
	return docStyle.Render(fmt.Sprintf(
		"Email:\n%s\n\nPassword:\n%s\n\n%s",
		l.emailInput.View(),
		l.passwordInput.View(),
		hint,
	))
}

func helpView() string {
	return helpStyle.Render("\ntab up/down • enter login • esc home")
}

// Ensure loginForm implements the model interface
var _ model = (*loginForm)(nil)
