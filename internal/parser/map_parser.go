package parser

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"stations-pathfinder/internal/domain"
)

const maxStations = 10000

// ParseMap parses the custom map format into a validated graph.
func ParseMap(path string) (domain.Graph, error) {
	f, err := os.Open(path)
	if err != nil {
		return domain.Graph{}, err
	}
	defer f.Close() // Ensure file is closed when done.

	g := domain.NewGraph()
	coords := make(map[string]struct{}) // Tracks used coordinates to enforce uniqueness across all stations.
	connSeen := make(map[string]struct{}) // Tracks seen connections to enforce undirected uniqueness.

	section := "" // Tracks the current section ("stations" or "connections") for parsing context.
	sc := bufio.NewScanner(f) // Reads the file line by line for efficient memory usage.

	for sc.Scan() {
		line := normalizeLine(sc.Text()) // Trim whitespace and ignore comments.
		if line == "" {
			continue
		}

		switch line {
		case "stations:":
			section = "stations"
			continue
		case "connections:":
			section = "connections"
			continue
		}

		switch section {
		case "stations":
			if err := parseStationLine(&g, coords, line); err != nil {
				return domain.Graph{}, err // Return on first error to prevent partial graph construction.
			}
		case "connections":
			if err := parseConnectionLine(&g, connSeen, line); err != nil {
				return domain.Graph{}, err // Return on first error to prevent partial graph construction.
			}
		default:
			return domain.Graph{}, errors.New("invalid file format")
		}
	}

	if err := sc.Err(); err != nil {
		return domain.Graph{}, err
	}

	return g, nil
}

// parseStationLine validates and inserts one station entry.
func parseStationLine(g *domain.Graph, coords map[string]struct{}, line string) error {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return errors.New("invalid station format")
	}
	name := strings.TrimSpace(parts[0])
	if !isValidStationName(name) {
		return domain.ErrInvalidStationName
	}
	if g.HasStation(name) {
		return domain.ErrDuplicateStationName
	}
	if len(g.Stations) >= maxStations {
		return domain.ErrTooManyStations
	}

	x, err := parseNonNegativeInt(parts[1])
	if err != nil {
		return domain.ErrInvalidCoordinate
	}
	y, err := parseNonNegativeInt(parts[2])
	if err != nil {
		return domain.ErrInvalidCoordinate
	}

	key := strconv.Itoa(x) + "," + strconv.Itoa(y) // Enforce global uniqueness of coordinates across all stations, regardless of name.
	
	if _, ok := coords[key]; ok {
		return domain.ErrDuplicateCoordinates
	}
	coords[key] = struct{}{}

	g.AddStation(domain.Station{Name: name, X: x, Y: y})
	return nil
}

// parseConnectionLine validates and inserts one undirected connection.
func parseConnectionLine(g *domain.Graph, seen map[string]struct{}, line string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return errors.New("invalid connection format")
	}

	a := strings.TrimSpace(parts[0])
	b := strings.TrimSpace(parts[1])

	if !g.HasStation(a) || !g.HasStation(b) {
		return domain.ErrUnknownStationInEdge
	}
// Create two keys for the undirected connection to check for duplicates in either direction.
	key1 := a + "|" + b
	key2 := b + "|" + a
	// Reject duplicates in either declared direction.
	if _, ok := seen[key1]; ok {
		return domain.ErrDuplicateConnection
	}
	if _, ok := seen[key2]; ok {
		return domain.ErrDuplicateConnection
	}
	seen[key1] = struct{}{}

	g.AddConnection(a, b)
	return nil
}

// parseNonNegativeInt parses an integer and enforces n >= 0.
func parseNonNegativeInt(s string) (int, error) {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil || n < 0 {
		return 0, errors.New("invalid non-negative integer")
	}
	return n, nil
}
