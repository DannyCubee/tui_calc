package main

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
	"os"
	"os/exec"
	"runtime"
)

type model struct {
	viewport viewport.Model
}

func whatOS() string {
	hostOS := runtime.GOOS

	return hostOS
}

var osMap = make(map[string]string)

func initApp() {
	osMap["linux"] = "clear"
	osMap["windows"] = "cls"
	osMap["darwin"] = "clear"

	cmd := exec.Command(osMap[whatOS()])
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getTerminalSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return width, height
	}
	return width, height
}

func main() {
	whatOS()
	initApp()
	width, height := getTerminalSize()
	m := model{
		viewport: viewport.Model{
			Width:  width - 3,
			Height: height - 2,
		},
	}
	p := tea.NewProgram(m)
	p.Run()
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case *tea.KeyMsg:
		switch msg.String() {
		case "cmd + q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {

	innerBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.viewport.Width / 2).
		Height(m.viewport.Height / 2).
		Padding(2).
		Render("Inner box")

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.viewport.Width).
		Height(m.viewport.Height).
		Render(innerBox)
}
