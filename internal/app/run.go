package app

import (
	"fmt"
	"io"

	"stations-pathfinder/internal/cli"
	"stations-pathfinder/internal/output"
	"stations-pathfinder/internal/parser"
	"stations-pathfinder/internal/routing"
	"stations-pathfinder/internal/scheduler"
)

// Run wires the CLI flow from arguments to formatted movement output.
func Run(args []string, stdout, stderr io.Writer) int {
	cfg, err := cli.ParseArgs(args) // ParseArgs validates and normalizes command line arguments into a Config struct.
	if err != nil {
		writeErr(stderr, err)
		return 1
	}
// The rest of the flow is to parse the map, build paths, plan moves, and write output.
	g, err := parser.ParseMap(cfg.MapPath)
	if err != nil {
		writeErr(stderr, err)
		return 1
	}
// BuildPathSet returns all paths from start to end, which the scheduler will use to plan moves.
	pathSet, err := routing.BuildPathSet(g, cfg.Start, cfg.End)
	if err != nil {
		writeErr(stderr, err)
		return 1
	}
// PlanMoves returns a list of turns, which the output package will format and write to stdout.
	turns, err := scheduler.PlanMoves(pathSet, cfg.TrainCount, cfg.Start, cfg.End)
	if err != nil {
		writeErr(stderr, err)
		return 1
	}
// WriteTurns writes the planned turns to stdout in the required format.
	if err := output.WriteTurns(stdout, turns); err != nil {
		writeErr(stderr, err)
		return 1
	}

	return 0
}

// writeErr keeps all error output in the required "Error: <message>" format.
func writeErr(w io.Writer, err error) {
	fmt.Fprintln(w, "Error:", err)
}
