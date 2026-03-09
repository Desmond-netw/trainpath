# Stations Pathfinder

## Project Overview
Stations Pathfinder is a Go command-line application that:
1. Parses a railway map file.
2. Finds a valid route from a start station to an end station (BFS shortest path).
3. Simulates train movements turn by turn.
4. Prints the movement plan to standard output.

The program accepts 4 positional arguments:
1. Map file path
2. Start station
3. End station
4. Number of trains

Output format:
1. Each move is `T<train_id>-<station>`.
2. One line per turn.
3. Multiple moves in the same turn are space-separated.

## Setup and Installation Instructions
### Prerequisites
1. Go installed (version supporting Go modules).

### Clone and enter the project
```bash
git clone <your-repo-url>
cd pathfinder
```

### Verify project builds and tests
```bash
go test ./...
```

## Usage Guide
### Basic command
```bash
go run ./cmd/stations-pathfinder <map-file> <start> <end> <train-count>
```

 1. ./cmd/stations-pathfinder - Project name
 2. map-file - Path to the network map to be use
 3. start - starting station for the journey
 4. end - name of the end station
 5. train-count - Total number of unique trains required for the mapping

### Example with a custom map
Create `example.map`:
```txt
stations:
a,1,1
b,2,1

connections:
a-b
```

Run:
```bash
go run ./cmd/stations-pathfinder example.map a b 1
```

Expected output:
```txt
T1-b
```

### Error behavior
On invalid input or map parsing failure, the app:
1. Prints `Error: <message>` to stderr.
2. Exits with status code `1`.

## Additional Features / Bonus Functionality
1. Strict map validation:
   - Station name validation (`[a-z0-9_]+`)
   - Duplicate station detection
   - Duplicate coordinate detection
   - Duplicate connection detection
   - Unknown station detection in connections
2. Turn-based occupancy rules:
   - Same edge cannot be used twice in one turn
   - Intermediate stations are single-occupancy per turn
   - Start/end stations allow multiple trains
3. Integration test coverage for app flow:
   - Successful run path
   - Error path (`start == end`)
