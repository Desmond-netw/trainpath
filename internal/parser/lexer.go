package parser

import "strings"

// normalizeLine strips inline comments and trims surrounding whitespace.
func normalizeLine(line string) string {
	if i := strings.Index(line, "#"); i >= 0 {
		line = line[:i]
	}
	return strings.TrimSpace(line)
}
