package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	v, err := StartViewInit()
	if err != nil {
		panic(err)
	}
	tea.NewProgram(v).Run()
}
