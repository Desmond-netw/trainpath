package output

import (
	"fmt"
	"io"
	"strings"

	"stations-pathfinder/internal/domain"
)

// WriteTurns prints movements as one line per turn (e.g. "T1-a T2-b").
func WriteTurns(w io.Writer, turns []domain.Turn) error {
	for _, turn := range turns {
		parts := make([]string, 0, len(turn))
		for _, mv := range turn {
			parts = append(parts, fmt.Sprintf("T%d-%s", mv.TrainID, mv.To))
		}
		if _, err := fmt.Fprintln(w, strings.Join(parts, " ")); err != nil {
			return err
		}
	}
	return nil
}
