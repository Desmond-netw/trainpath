package main

import (
	"os"

	"stations-pathfinder/internal/app"
)

// main is the CLI entrypoint and delegates all work to internal app orchestration.
func main() {
	os.Exit(app.Run(os.Args, os.Stdout, os.Stderr))
}
