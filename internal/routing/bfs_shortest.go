package routing

import "stations-pathfinder/internal/domain"

// ShortestPathBFS returns one shortest unweighted path between start and end.
func ShortestPathBFS(g domain.Graph, start, end string) (domain.Path, error) {
	if !g.HasStation(start) {
		return nil, domain.ErrStartStationNotFound
	}
	if !g.HasStation(end) {
		return nil, domain.ErrEndStationNotFound
	}
// Standard BFS implementation using a queue, visited set, and predecessor map.
	q := []string{start}
	visited := map[string]bool{start: true}
	prev := map[string]string{}
	
// Loop until the queue is empty, exploring neighbors level by level.
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur == end {
			return rebuildPath(prev, start, end), nil
		}

		for _, nb := range g.Neighbors(cur) {
			if visited[nb] {
				continue
			}
			visited[nb] = true
			prev[nb] = cur
			q = append(q, nb)
		}
	}

	return nil, domain.ErrNoPath
}

// rebuildPath reconstructs a path from BFS predecessor links.
func rebuildPath(prev map[string]string, start, end string) domain.Path {
	path := domain.Path{end}
	for cur := end; cur != start; {
		cur = prev[cur]
		path = append(path, cur)
	}
	reverse(path)
	return path
}

// reverse reverses a slice in place.
func reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
