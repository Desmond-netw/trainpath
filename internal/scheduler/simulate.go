package scheduler

import (
	"sort"

	"stations-pathfinder/internal/domain"
)

// trainState tracks one train along its assigned path.
type trainState struct {
	id   int
	path domain.Path
	pos  int
	done bool
}

// simulateTurns runs turn-based simulation with potentially different paths per train.
func simulateTurns(paths []domain.Path, start, end string) []domain.Turn {
	trains := make([]trainState, len(paths))
	for i := range paths {
		trains[i] = trainState{id: i + 1, path: paths[i], pos: 0}
	}

	var turns []domain.Turn

	for {
		turnOcc := NewOccupancy()
		turn := make(domain.Turn, 0)
		allDone := true

		order := make([]int, 0, len(trains))
		for i := range trains {
			order = append(order, i)
		}
		sort.Slice(order, func(i, j int) bool {
			a := trains[order[i]]
			b := trains[order[j]]
			// Prioritize trains already in transit and those closer to destination.
			aStarted := a.pos > 0
			bStarted := b.pos > 0
			if aStarted != bStarted {
				return aStarted
			}
			aRem := len(a.path) - 1 - a.pos
			bRem := len(b.path) - 1 - b.pos
			if aRem != bRem {
				return aRem < bRem
			}
			if len(a.path) != len(b.path) {
				return len(a.path) < len(b.path)
			}
			return a.id < b.id
		})

		for _, idx := range order {
			t := &trains[idx]
			if t.done {
				continue
			}
			allDone = false

			if t.pos+1 >= len(t.path) {
				t.done = true
				continue
			}

			from := t.path[t.pos]
			to := t.path[t.pos+1]
			ek := edgeKey(from, to)
			if turnOcc.Edge[ek] {
				continue
			}
			// Start/end can hold unlimited trains; intermediate stations cannot.
			if to != end && turnOcc.Station[to] {
				continue
			}
			turnOcc.Edge[ek] = true

			if to != end && to != start {
				turnOcc.Station[to] = true
			}

			t.pos++
			if t.pos == len(t.path)-1 {
				t.done = true
			}

			turn = append(turn, domain.Move{TrainID: t.id, To: to})
		}

		if allDone {
			break
		}
		if len(turn) > 0 {
			turns = append(turns, turn)
		}
	}

	return turns
}
