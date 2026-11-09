package main

import (
	"fmt"
	"os"

	"github.com/vaibe-studio/vaibe-os-installer/internal/installer"
)

const usage = `vaibe-os — the operating layer for your AI agents (installer)

Usage:
  vaibe-os init <folder> [--name <name>]   create a new instance workspace
  vaibe-os version                          print version
  vaibe-os validate                         (planned: validate instance structure)
  vaibe-os doctor                           (planned: runtime invariants)
  vaibe-os connect <instance>               (planned: link to another instance)
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "init":
		if err := installer.RunInit(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "vaibe-os init: %v\n", err)
			os.Exit(1)
		}
	case "version":
		fmt.Println("vaibe-os", installer.Version)
	case "validate", "doctor", "connect":
		fmt.Printf("vaibe-os %s: not implemented yet (planned for v0.1.x)\n", os.Args[1])
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n%s", os.Args[1], usage)
		os.Exit(1)
	}
}
