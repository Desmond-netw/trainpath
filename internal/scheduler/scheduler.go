package scheduler

import (
	"errors"

	"stations-pathfinder/internal/domain"
)

// Planner describes a schedule builder from path set to turn-by-turn moves.
type Planner interface {
	PlanMoves(pathSet []domain.Path, trainCount int, start, end string) ([]domain.Turn, error)
}

// PlanMoves validates inputs and applies the current scheduling strategy.
func PlanMoves(pathSet []domain.Path, trainCount int, start, end string) ([]domain.Turn, error) {
	if len(pathSet) == 0 {
		return nil, domain.ErrNoPath
	}
	if trainCount <= 0 {
		return nil, errors.New("invalid train count")
	}

	assignments := assignPaths(pathSet, trainCount)
	return simulateTurns(assignments, start, end), nil
}
