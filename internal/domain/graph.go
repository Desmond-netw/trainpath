package domain

// Graph stores stations and undirected adjacency using set-like maps.
type Graph struct {
	Stations map[string]Station
	Adj      map[string]map[string]struct{}
}

// NewGraph creates an empty graph with initialized maps.
func NewGraph() Graph {
	return Graph{
		Stations: make(map[string]Station),
		Adj:      make(map[string]map[string]struct{}),
	}
}

// AddStation inserts or replaces a station and ensures an adjacency bucket exists.
func (g *Graph) AddStation(s Station) {
	g.Stations[s.Name] = s
	if _, ok := g.Adj[s.Name]; !ok {
		g.Adj[s.Name] = make(map[string]struct{})
	}
}

// HasStation reports whether a station name exists in the graph.
func (g *Graph) HasStation(name string) bool {
	_, ok := g.Stations[name]
	return ok
}

// AddConnection inserts an undirected edge between two stations.
func (g *Graph) AddConnection(a, b string) {
	if _, ok := g.Adj[a]; !ok {
		g.Adj[a] = make(map[string]struct{})
	}
	if _, ok := g.Adj[b]; !ok {
		g.Adj[b] = make(map[string]struct{})
	}
	g.Adj[a][b] = struct{}{} 
	g.Adj[b][a] = struct{}{}
}

// Neighbors returns adjacent station names for a station.
func (g *Graph) Neighbors(name string) []string {
	n := make([]string, 0, len(g.Adj[name]))
	for v := range g.Adj[name] {
		n = append(n, v)
	}
	return n
}
