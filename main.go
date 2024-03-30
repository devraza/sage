package main

import (
	"fmt"
	"os"
  "strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/prometheus-community/pro-bing"
)

var (
	checkMark = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	docStyle  = lipgloss.NewStyle().Margin(1, 2)
  titleStyle = lipgloss.NewStyle().Bold(true)
)

type keyMap struct {
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit},
	}
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type model struct {
	keys      keyMap
	help      help.Model
	addresses []string
}

func initialModel() model {
	return model{
		keys:      keys,
		help:      help.New(),
		addresses: []string{"100.64.0.1", "100.64.0.2"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("Pinging hosts...") + strings.Repeat("\n", 2)
	for _, address := range m.addresses {
		_, err := probing.NewPinger(fmt.Sprintf("%s", address))
		if err != nil {
			panic(err)
		} else {
			s += fmt.Sprintf("%s  %s\n", checkMark, address)
		}
	}

	helpView := m.help.View(m.keys)
	return docStyle.Render(s + "\n" + helpView)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
