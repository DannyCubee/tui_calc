package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"os/exec"
	"runtime"
)

func whatOS() string {
	hostOS := runtime.GOOS
	fmt.Println(hostOS)

	return hostOS
}

var osMap = make(map[string]string)

func initApp() {
	osMap["linux"] = "clear"
	osMap["windows"] = "cls"

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
	fmt.Printf("Terminal size: %dx%d\n", width, height)
}
