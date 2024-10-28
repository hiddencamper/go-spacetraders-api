package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	APITools "github.com/hiddencamper/go-spacetraders-api/APITools"
	wrap "github.com/hiddencamper/go-wordwrap"
)

type StartView struct {
	GetStatus APITools.GetStatus
	height    int
	width     int
	cursor    int
	options   []string
}

func StartViewInit() (StartView, error) {
	g, err := APITools.API_GetStatus()
	var s StartView
	s.GetStatus = *g
	s.cursor = 0
	s.options = []string{"Announcements", "Leaderboards", "Quit"}
	s.height = 0
	s.width = 0
	return s, err
}

func (m StartView) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m StartView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter":
			if m.options[m.cursor] == "Quit" {
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		if m.width > 120 {
			m.width = 120
		}
	}
	return m, nil
}

func (m StartView) View() string {
	if m.height == 0 || m.width == 0 {
		//TODO: Add a timer or delay here so that execution can continue
		//after a delay with a default size of 80w 24/h
		return ""
	}
	g := &m.GetStatus
	s := "Space Traders Terminal User Interface\n\n"
	s += fmt.Sprintf("Space Traders API Version: %s\n", g.Version)
	s += fmt.Sprintf("Server Status: %s\n", g.Status)
	s += fmt.Sprintf("Last Reset Date: %s\n", g.ResetDate)
	s += fmt.Sprintf("\n%s\n\n", wrap.WordWrap(g.Description, m.width))

	for i, o := range m.options {
		if i == m.cursor {
			s += " > "
		} else {
			s += "   "
		}
		s += fmt.Sprintf("%v:%s\n", i+1, o)
	}
	s += "\n\nPress q or Control+C to quit\n"
	return s
}
