package domain

import "errors" // Package domain contains core data structures and canonical errors used across the application.

// Canonical domain-level errors used across parser, routing, and scheduling.
var (
	ErrNoPath                 = errors.New("no path between the start and end stations")
	ErrStartStationNotFound   = errors.New("start station does not exist")
	ErrEndStationNotFound     = errors.New("end station does not exist")
	ErrTooManyStations        = errors.New("map contains more than 10000 stations")
	ErrDuplicateStationName   = errors.New("duplicate station names")
	ErrInvalidStationName     = errors.New("invalid station names")
	ErrInvalidCoordinate      = errors.New("coordinates must be valid non-negative integers")
	ErrDuplicateCoordinates   = errors.New("two stations exist at the exact same coordinate location")
	ErrDuplicateConnection    = errors.New("duplicate connections")
	ErrUnknownStationInEdge   = errors.New("connection references a station which does not exist")
	ErrNotImplementedStrategy = errors.New("not implemented")
)
