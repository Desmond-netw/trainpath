package scheduler

import "stations-pathfinder/internal/domain"

const maxSchedulingCandidates = 8

// assignPaths balances trains over candidate paths with a small local optimizer.
func assignPaths(pathSet []domain.Path, trainCount int) []domain.Path {
	if len(pathSet) > maxSchedulingCandidates {
		pathSet = pathSet[:maxSchedulingCandidates]
	}
	start, end := pathSet[0][0], pathSet[0][len(pathSet[0])-1]

	// Greedy seed: pick the path that minimizes turns for trains assigned so far.
	assignments := make([]domain.Path, trainCount)
	for i := 0; i < trainCount; i++ {
		best, bestTurns := 0, 0
		for p := range pathSet {
			assignments[i] = pathSet[p]
			t := len(simulateTurns(assignments[:i+1], start, end))
			if p == 0 || t < bestTurns || (t == bestTurns && len(pathSet[p]) < len(pathSet[best])) {
				best, bestTurns = p, t
			}
		}
		assignments[i] = pathSet[best]
	}

	// Local improvement: single-train reassignment if total turns improve.
	bestTurns := len(simulateTurns(assignments, start, end))
	for changed := true; changed; {
		changed = false
		for i := range assignments {
			orig := assignments[i]
			for p := range pathSet {
				if pathsEqual(pathSet[p], orig) {
					continue
				}
				assignments[i] = pathSet[p]
				t := len(simulateTurns(assignments, start, end))
				if t < bestTurns {
					bestTurns = t
					changed = true
					orig = assignments[i]
				}
			}
			assignments[i] = orig
		}
	}

	// Small exact search over path counts (if tractable) to avoid greedy dead-ends.
	if len(pathSet) <= 8 && choose(trainCount+len(pathSet)-1, len(pathSet)-1) <= 100000 {
		counts := make([]int, len(pathSet))
		var search func(int, int)
		search = func(i, left int) {
			if i == len(pathSet)-1 {
				counts[i] = left
				cand := make([]domain.Path, 0, trainCount)
				rem := append([]int(nil), counts...)
				for len(cand) < trainCount {
					for p := range pathSet {
						if rem[p] > 0 {
							cand = append(cand, pathSet[p])
							rem[p]--
						}
					}
				}
				if t := len(simulateTurns(cand, start, end)); t < bestTurns {
					bestTurns = t
					copy(assignments, cand)
				}
				return
			}
			for c := 0; c <= left; c++ {
				counts[i] = c
				search(i+1, left-c)
			}
		}
		search(0, trainCount)
	}

	return assignments
}

func pathsEqual(a, b domain.Path) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func choose(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	r := 1
	for i := 1; i <= k; i++ {
		r = r * (n - k + i) / i
	}
	return r
}
