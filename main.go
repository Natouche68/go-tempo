package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mbndr/figlet4go"

	"github.com/Natouche68/go-tempo/styles"
	timeutils "github.com/Natouche68/go-tempo/time-utils"
)

type Model struct {
	timeLeft     int
	workTime     int
	pauseTime    int
	currentPhase string
	ticking      bool
}

type TickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return doTick()
}

var ascii = figlet4go.NewAsciiRender()

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case " ":
			m.ticking = !m.ticking

		case "r":
			if m.currentPhase == "work" {
				m.timeLeft = m.workTime
			} else if m.currentPhase == "pause" {
				m.timeLeft = m.pauseTime
			}
		}

	case TickMsg:
		if m.ticking {
			m.timeLeft--

			if m.timeLeft < 0 {
				if m.currentPhase == "work" {
					m.currentPhase = "pause"
					m.timeLeft = m.pauseTime
				} else if m.currentPhase == "pause" {
					m.currentPhase = "work"
					m.timeLeft = m.workTime
				}
			}
		}
		return m, doTick()
	}

	return m, nil
}

func (m Model) View() string {
	var s string

	title := styles.TitleStyle.Render("Go Tempo")

	clock, err := ascii.Render(timeutils.SecondToClock(m.timeLeft))
	if err != nil {
		return styles.ErrorStyle.Render("An error occured : \n" + err.Error())
	}
	clock = strings.TrimSuffix(clock, "\n")

	var label string
	switch m.currentPhase {
	case "work":
		if m.ticking {
			label = styles.LabelStyle.
				Copy().
				Foreground(styles.WorkColor).
				Render("- Work -")
		} else {
			label = styles.LabelStyle.
				Copy().
				Foreground(styles.WorkColor).
				Italic(true).
				Render("- Work (paused) -")
		}
	case "pause":
		if m.ticking {
			label = styles.LabelStyle.
				Copy().
				Foreground(styles.PauseColor).
				Render("- Pause -")
		} else {
			label = styles.LabelStyle.
				Copy().
				Foreground(styles.PauseColor).
				Italic(true).
				Render("- Pause (paused) -")
		}
	}

	help := styles.HelpStyle.Render("q : quit  •  space : pause/resume  •  r : reset")

	s = lipgloss.JoinVertical(lipgloss.Center, title, styles.ClockStyle.Render(clock), label, help)

	return s
}

func main() {
	p := tea.NewProgram(Model{
		timeLeft:     25 * 60,
		workTime:     25 * 60,
		pauseTime:    5 * 60,
		currentPhase: "work",
		ticking:      true,
	}, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println(styles.ErrorStyle.Render("There's been an error : " + err.Error()))
		os.Exit(1)
	}
}
