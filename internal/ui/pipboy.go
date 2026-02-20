package ui

import (
	"fmt"
	"strings"
	"time"

	"atlas.facade/internal/server"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	amber      = lipgloss.Color("#FFB642")
	rusty      = lipgloss.Color("#5E4737")
	ncrRed     = lipgloss.Color("#FF0000")
	onyx       = lipgloss.Color("#050505")

	screenStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(amber).
			Background(onyx)

	headerStyle = lipgloss.NewStyle().
			Foreground(onyx).
			Background(amber).
			Padding(0, 1).
			Bold(true)

	logStyle = lipgloss.NewStyle().Foreground(amber)
	dimStyle = lipgloss.NewStyle().Foreground(rusty)
)

type Model struct {
	server  *server.Server
	logs    []server.LogEntry
	width   int
	height  int
}

func NewModel(s *server.Server) Model {
	return Model{
		server: s,
		logs:   []server.LogEntry{},
	}
}

func (m Model) Init() tea.Cmd {
	return waitForLog(m.server.LogChan)
}

type logMsg server.LogEntry

func waitForLog(ch chan server.LogEntry) tea.Cmd {
	return func() tea.Msg {
		return logMsg(<-ch)
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case logMsg:
		m.logs = append(m.logs, server.LogEntry(msg))
		if len(m.logs) > 10 {
			m.logs = m.logs[1:]
		}
		return m, waitForLog(m.server.LogChan)
	}
	return m, nil
}

func (m Model) View() string {
	var sb strings.Builder

	header := headerStyle.Render(" ATLAS PIP-BOY 3000 ") + " " + logStyle.Render(fmt.Sprintf("PORT :%d", m.server.Port))
	sb.WriteString(header + "\n\n")

	sb.WriteString(logStyle.Render("SYSTEM STATUS: NOMINAL") + "\n")
	sb.WriteString(logStyle.Render("V.A.T.S. PROXY ENGAGED") + "\n\n")

	sb.WriteString(dimStyle.Render("ACTIVE BLUEPRINTS:") + "\n")
	for _, r := range m.server.Routes {
		sb.WriteString(logStyle.Render(fmt.Sprintf(" > %-6s %s", r.Method, r.Path)) + "\n")
	}
	sb.WriteString("\n")

	sb.WriteString(dimStyle.Render("REQUEST FEED:") + "\n")
	if len(m.logs) == 0 {
		sb.WriteString(dimStyle.Render(" [WAITING FOR TRAFFIC...]") + "\n")
	} else {
		for _, l := range m.logs {
			ts := l.Timestamp.Format("15:04:05")
			line := fmt.Sprintf("[%s] %-6s %-15s %d (%v)", ts, l.Method, l.Path, l.Status, l.Duration.Round(time.Millisecond))
			sb.WriteString(logStyle.Render(line) + "\n")
		}
	}

	sb.WriteString("\n" + dimStyle.Render("PRESS 'Q' TO SHUT DOWN"))

	return screenStyle.Render(sb.String())
}
