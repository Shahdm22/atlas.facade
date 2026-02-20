package main

import (
	"flag"
	"fmt"
	"os"

	"atlas.facade/internal/model"
	"atlas.facade/internal/server"
	"atlas.facade/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fezcode/go-piml"
)

var Version = "dev"

func main() {
	port := flag.Int("port", 4000, "Port to run the mock server on")
	blueprintPath := flag.String("file", "routes.piml", "PIML file containing route definitions")
	versionFlag := flag.Bool("version", false, "Show version information")
	flag.BoolVar(versionFlag, "v", false, "Show version information")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("atlas.facade v%s\n", Version)
		return
	}

	// 1. Load Blueprint
	data, err := os.ReadFile(*blueprintPath)
	if err != nil {
		fmt.Printf("Error: Blueprint file '%s' not found.\n", *blueprintPath)
		os.Exit(1)
	}

	var blueprint model.Blueprint
	if err := piml.Unmarshal(data, &blueprint); err != nil {
		fmt.Printf("Error parsing blueprint: %v\n", err)
		os.Exit(1)
	}

	// 2. Init Server
	srv := server.NewServer(*port, blueprint.Routes)

	// 3. Start Server Goroutine
	go func() {
		if err := srv.Start(); err != nil {
			fmt.Printf("Server Error: %v\n", err)
			os.Exit(1)
		}
	}()

	// 4. Start TUI
	p := tea.NewProgram(ui.NewModel(srv), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("TUI Error: %v\n", err)
		os.Exit(1)
	}
}
