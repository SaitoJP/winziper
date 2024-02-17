package textinput

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type model struct {
	err       error
	textInput textinput.Model
}

func Run(placeholder string) string {
	p := tea.NewProgram(initialModel(placeholder))
	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	m2, ok := m.(model)
	if !ok {
		log.Fatal("キャスト失敗")
	}

	return m2.textInput.Value()
}

func initialModel(placeholder string) model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 60

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf("\n%s\n\n%s", m.textInput.View(), "(esc to quit)") + "\n"
}
