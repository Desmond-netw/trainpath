package routing

import "stations-pathfinder/internal/domain"

// PathFinder describes a routing strategy that builds a path set for scheduling.
type PathFinder interface {
	BuildPathSet(g domain.Graph, start, end string) ([]domain.Path, error)
}

// BuildPathSet selects the set of candidate paths for the scheduler.
func BuildPathSet(g domain.Graph, start, end string) ([]domain.Path, error) {
	return FindCandidatePaths(g, start, end)
}
