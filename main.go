package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	timeLeft  int
	workTime  int
	pauseTime int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}

	return m, nil
}

func (m Model) View() string {
	s := "Go Tempo"

	return s
}

func main() {
	p := tea.NewProgram(Model{
		timeLeft:  25 * 60,
		workTime:  25 * 60,
		pauseTime: 5 * 60,
	}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
