package scheduler

// Occupancy tracks resources consumed within a single turn.
type Occupancy struct {
	Station map[string]bool
	Edge    map[string]bool
}

// NewOccupancy creates empty per-turn occupancy sets.
func NewOccupancy() Occupancy {
	return Occupancy{
		Station: make(map[string]bool),
		Edge:    make(map[string]bool),
	}
}

// edgeKey canonicalizes an undirected edge into a stable key.
func edgeKey(a, b string) string {
	if a < b {
		return a + "|" + b
	}
	return b + "|" + a
}
