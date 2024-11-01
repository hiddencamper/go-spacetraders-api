package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/viewport"
	APITools "github.com/hiddencamper/go-spacetraders-api/APITools"
	wrap "github.com/hiddencamper/go-wordwrap"
)

type StartView struct {
	GetStatus APITools.GetStatus
	height    int
	width     int
	cursor    int
	options   []string
	current   string
	subnum    int
	view      viewport.Model
}

func StartViewInit() (StartView, error) {
	g, err := APITools.API_GetStatus()
	var s StartView
	s.GetStatus = *g
	s.cursor = 0
	s.options = []string{"Description", "Announcements", "Links", "Leaderboards", "Quit"}
	s.height = 0
	s.width = 0
	s.current = s.options[0]
	s.subnum = 0
	s.view = viewport.New(120, 10)
	return s, err
}

type TickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m StartView) Init() tea.Cmd {
	return tea.Batch(tea.ClearScreen, doTick())
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
			m.subnum = 0
			m.current = m.options[m.cursor]
		case "left":
			if m.subnum > 0 {
				m.subnum--
			}
		case "right":
			if m.current == "Announcements" && m.subnum < len(m.GetStatus.Announcements)-1 {
				m.subnum++
			}
			if m.current == "Links" && m.subnum < len(m.GetStatus.Links)-1 {
				m.subnum++
			}
		}
	case TickMsg:
		if m.width == 0 || m.height == 0 {
			m.height = 24
			m.width = 80
		}
		return m, nil
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.view.Width = m.width
		if m.width > 120 {
			m.width = 120
		}
	}
	return m, nil
}

func (m StartView) View() string {
	if m.height == 0 || m.width == 0 {
		return ""
	}
	g := &m.GetStatus
	s := "Space Traders Terminal User Interface\n\n"
	s += fmt.Sprintf("Space Traders API Version: %s\n", g.Version)
	s += fmt.Sprintf("Server Status: %s\n", g.Status)
	s += fmt.Sprintf("Last Reset Date: %s\n", g.ResetDate)
	t := "\n"
	if m.current == "Description" {
		t += fmt.Sprintf("\n\n%s\n\n", wrap.WordWrap(g.Description, m.width))
	}
	if m.current == "Announcements" {
		if len(g.Announcements) <= 0 {
			t += "\nNo announcements\n"
		} else {

			t += fmt.Sprintf("\nAnnouncement %d of %d:  %s\n", m.subnum+1, len(g.Announcements), g.Announcements[m.subnum].Title)
			t += fmt.Sprintf("%s\n\n", wrap.WordWrap(g.Announcements[m.subnum].Body, m.width))
		}
	}
	if m.current == "Leaderboards" {

		t += "\nLeaderboards not implemented yet\n"
	}
	if m.current == "Links" {
		if len(g.Links) <= 0 {
			t += "\nNo links\n"
		} else {
			t += fmt.Sprintf("\nLink %d of %d:  %s\n", m.subnum+1, len(g.Links), g.Links[m.subnum].Name)
			t += fmt.Sprintf("%s\n\n", wrap.WordWrap(g.Links[m.subnum].URL, m.width))
		}
	}
	m.view.SetContent(t)
	u := "\n"
	for i, o := range m.options {
		if i == m.cursor {
			u += " > "
		} else {
			u += "   "
		}
		u += fmt.Sprintf("%v:%s\n", i+1, o)
	}
	u += "\n\nPress q or Control+C to quit\n"
	u += "Press left or right to change subpage\n"
	return s + m.view.View() + u
}
