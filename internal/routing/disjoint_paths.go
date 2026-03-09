package routing

import (
	"sort"
	"strings"

	"stations-pathfinder/internal/domain"
)

const maxCandidatePaths = 64

// FindCandidatePaths returns multiple simple path candidates sorted by length.
func FindCandidatePaths(g domain.Graph, start, end string) ([]domain.Path, error) {
	shortest, err := ShortestPathBFS(g, start, end)
	if err != nil {
		return nil, err
	}
// Enumerate simple paths up to a limit, then sort and deduplicate them.
	candidates := enumerateSimplePaths(g, start, end, maxCandidatePaths)
	if len(candidates) == 0 {
		return []domain.Path{shortest}, nil
	}
// Sort by length, then lexicographically to ensure deterministic order.
	sort.Slice(candidates, func(i, j int) bool {
		if len(candidates[i]) != len(candidates[j]) {
			return len(candidates[i]) < len(candidates[j])
		}
		return pathKey(candidates[i]) < pathKey(candidates[j])
	})

	// Deduplicate paths using a map keyed by their string representation.
	seen := make(map[string]struct{}, len(candidates))
	unique := make([]domain.Path, 0, len(candidates))
	for _, p := range candidates {
		k := pathKey(p)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		unique = append(unique, p)
	}
	return unique, nil
}
// enumerateSimplePaths performs a depth-limited DFS to find simple paths from start to end.
func enumerateSimplePaths(g domain.Graph, start, end string, limit int) []domain.Path {
	if limit <= 0 {
		return nil
	}

	maxDepth := len(g.Stations)
	visited := make(map[string]bool, len(g.Stations))
	visited[start] = true

	cur := domain.Path{start}
	out := make([]domain.Path, 0, limit)

	var dfs func(node string)
	dfs = func(node string) {
		if len(out) >= limit {
			return
		}
		if node == end {
			p := make(domain.Path, len(cur))
			copy(p, cur)
			out = append(out, p)
			return
		}
		if len(cur) >= maxDepth {
			return
		}

		nbrs := g.Neighbors(node)
		sort.Strings(nbrs)
		for _, n := range nbrs {
			if visited[n] {
				continue
			}
			visited[n] = true
			cur = append(cur, n)
			dfs(n)
			cur = cur[:len(cur)-1]
			visited[n] = false
			if len(out) >= limit {
				return
			}
		}
	}

	dfs(start)
	return out
}
// pathKey generates a string key for a path by joining station names with "->". This is used for sorting and deduplication.
func pathKey(p domain.Path) string {
	return strings.Join(p, "->")
}
