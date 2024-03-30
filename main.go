package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"

    "github.com/prometheus-community/pro-bing"
)

var (
	checkMark = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
)

type model struct {
  addresses  []string
}

func initialModel() model {
	return model{
    addresses:  []string{"100.64.0.1", "100.64.0.2"},
	}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    s := "Pinging hosts...\n\n"
    for _, address := range m.addresses {
        _, err := probing.NewPinger(fmt.Sprintf("%s", address))
        if err != nil {
	        panic(err)
        } else {
          s += fmt.Sprintf("%s  %s\n", checkMark, address)
        }
    }
    s += "\nPress q to quit.\n"
    return s
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v", err)
        os.Exit(1)
    }
}
