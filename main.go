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
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Printf("atlas.facade v%s\n", Version)
		return
	}
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help") {
		fmt.Println("Atlas Facade - Retro-future Pip-Boy style mock API server.")
		fmt.Println("\nUsage:")
		fmt.Println("  atlas.facade [options]")
		fmt.Println("\nOptions:")
		fmt.Println("  -port int     Port to run the mock server on (default 4000)")
		fmt.Println("  -file string  PIML file containing route definitions (default \"routes.piml\")")
		fmt.Println("  -v, -version  Show version information")
		fmt.Println("  -h, -help     Show this help")
		return
	}

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
