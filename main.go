package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
	"os"
	"os/exec"
	"runtime"
)

type Button struct {
	ID      string
	Pressed bool
	Style   lipgloss.Style
	Color   string
}

type model struct {
	viewport viewport.Model
	buttons  map[string]*Button
}

func main() {
	whatOS()
	width, height := getTerminalSize()
	m := model{
		viewport: viewport.Model{
			Width:  width / 3,
			Height: height - 2,
		},
		buttons: map[string]*Button{
			"6": &Button{
				ID:    "6",
				Style: lipgloss.NewStyle(),
				Color: "green",
			},
		},
	}

	initApp()
	p := tea.NewProgram(m)
	p.Run()
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC || msg.String() == "q" || msg.String() == "Q" {
			return m, tea.Quit
		}
		if msg.String() == "6" {
			button := m.buttons["6"]
			button.Pressed = !button.Pressed
			if button.Pressed {
				button.Color = "red"
			} else {
				button.Color = "green"
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	button := m.buttons["6"]
	buttonStyle := button.Style.Border(lipgloss.RoundedBorder())
	buttonView := buttonStyle.Render(button.ID)

	innerBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.viewport.Width - 2).
		Height(m.viewport.Height / 4).
		Padding(2).
		Render("Inner box")

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.viewport.Width).
		Height(m.viewport.Height).
		Render(innerBox + "\n" + buttonView)
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
	fmt.Println("running " + osMap[whatOS()])
	cmd.Run()
}

func getTerminalSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return width, height
	}
	return width, height
}
