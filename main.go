package main

import (
  "fmt"
  "log"
  "os/exec"
  "strings"

  tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	address             = lipgloss.NewStyle().Bold(true)
	crossMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("210")).SetString("✗")
  checkMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("120")).SetString("✓")
  port                = lipgloss.NewStyle().Background(lipgloss.Color("236"))
  items               = []string{"443", "80", "2222"}
)

func initialModel() model {
	return model{
	  outputs:   make([]int, 0),
		err:       nil,
	}
}

type model struct {
  outputs   []int
	err       error
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
  return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
  }

  for _, i := range items {
    cmd := exec.Command("nc", "-zv", "devraza.giize.com", i, "-w", "1")
    out, _ := cmd.CombinedOutput()

    if len(strings.Split(string(out), " ")) < 12 {
      m.outputs = append(m.outputs, 1)
    } else {
      m.outputs = append(m.outputs, 0)
    }
  }
  return m, nil
}

func (m model) View() string {
  var combined string
  for idx, i := range m.outputs {
    if i == 1 {
      combined += fmt.Sprintf("%v  %v\n", checkMark.Render(), port.Render(" " + items[idx] + " "))
    } else if i == 0 {
      combined += fmt.Sprintf("%v  %v\n", crossMark.Render(), port.Render(" " + items[idx] + " "))
    }
  }

	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	} else {
		return fmt.Sprintf("%v\n%v", address.Render("devraza.giize.com"), combined)
	}
}
