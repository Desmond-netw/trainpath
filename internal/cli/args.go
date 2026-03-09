package cli

import (
	"errors"
	"strconv"
)

// Config contains normalized command line arguments.
type Config struct {
	MapPath    string
	Start      string
	End        string
	TrainCount int
}

// ParseArgs validates required positional arguments and returns typed config.
func ParseArgs(args []string) (Config, error) {
	if len(args) != 5 {
		return Config{}, errors.New("incorrect number of command line arguments")
	}
// Parse the number of trains and validate that it's a positive integer.
	n, err := strconv.Atoi(args[4])
	if err != nil || n <= 0 {
		return Config{}, errors.New("number of trains is not a valid positive integer")
	}
// Validate that the start and end stations are not the same.
	if args[2] == args[3] {
		return Config{}, errors.New("start and end station are the same")
	}
// If all validations pass, return the populated Config struct.
	return Config{
		MapPath:    args[1],
		Start:      args[2],
		End:        args[3],
		TrainCount: n,
	}, nil
}
