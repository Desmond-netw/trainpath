package parser

import "regexp"

// Station names may only contain lowercase letters, digits, and underscores.
var stationNameRE = regexp.MustCompile(`^[a-z0-9_]+$`)

// isValidStationName checks raw station name syntax.
func isValidStationName(s string) bool {
	return stationNameRE.MatchString(s)
}
