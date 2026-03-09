package domain

// Path is an ordered list of station names from start to end.
type Path []string

// Move is one train movement during a turn.
type Move struct {
	TrainID int
	To      string
}

// Turn contains all train moves executed in the same time step.
type Turn []Move
