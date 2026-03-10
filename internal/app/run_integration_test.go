package app_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"stations-pathfinder/internal/app"
)
// TestRun_MinimalNetwork_OneTrain tests the happy path of a single train moving from waterloo to euston.
func TestRun_MinimalNetwork_OneTrain(t *testing.T) {
	args := []string{
		"stations-pathfinder",
		"../../network_map/london.map",
		"waterloo",
		"euston",
		"1",
	}

	// Use bytes.Buffer to capture stdout and stderr for assertions.
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := app.Run(args, &stdout, &stderr)

	if code != 0 {
		t.Fatalf("expected exit code 0, got %d, stderr=%q", code, stderr.String())
	}

	// Assert that the output matches the expected format for one train moving from waterloo to euston.
	got := stdout.String()
	want := "T1-euston\n"
	if got != want {
		t.Fatalf("unexpected stdout\nwant: %q\ngot:  %q", want, got)
	}
// Assert that there are no errors on stderr.
	if stderr.Len() != 0 {
		t.Fatalf("expected empty stderr, got %q", stderr.String())
	}
}
// TestRun_ErrorWhenStartEqualsEnd tests that the application returns an error when the start and end stations are the same.
func TestRun_ErrorWhenStartEqualsEnd(t *testing.T) {
	args := []string{
		"stations-pathfinder",
		"../../network_map/london.map",
		"waterloo",
		"waterloo",
		"1",
	}
// Use bytes.Buffer to capture stdout and stderr for assertions.
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := app.Run(args, &stdout, &stderr)

	if code == 0 {
		t.Fatalf("expected non-zero exit code, got %d", code)
	}

	if stdout.Len() != 0 {
		t.Fatalf("expected empty stdout, got %q", stdout.String())
	}

	if stderr.Len() == 0 {
		t.Fatalf("expected error on stderr, got empty")
	}

	if got := stderr.String(); !strings.HasPrefix(got, "Error:") {
		t.Fatalf("expected stderr to start with %q, got %q", "Error:", got)
	}
}

func TestRun_ErrorWhenMapHasMoreThan10000Stations(t *testing.T) {
	var b strings.Builder
	b.WriteString("stations:\n")
	for i := 0; i <= 10000; i++ {
		fmt.Fprintf(&b, "s%d,%d,0\n", i, i)
	}
	b.WriteString("\nconnections:\n")

	mapPath := filepath.Join(t.TempDir(), "too_many_stations.map")
	if err := os.WriteFile(mapPath, []byte(b.String()), 0o644); err != nil {
		t.Fatalf("failed to write map file: %v", err)
	}

	args := []string{
		"stations-pathfinder",
		mapPath,
		"s0",
		"s1",
		"1",
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := app.Run(args, &stdout, &stderr)
	if code == 0 {
		t.Fatalf("expected non-zero exit code, got %d", code)
	}
	if stdout.Len() != 0 {
		t.Fatalf("expected empty stdout, got %q", stdout.String())
	}
	if got := stderr.String(); !strings.HasPrefix(got, "Error:") {
		t.Fatalf("expected stderr to start with %q, got %q", "Error:", got)
	}
}

func TestRun_BeginningAndTerminus_20Trains_CompletesWithin11Turns(t *testing.T) {
	args := []string{
		"stations-pathfinder",
		"../../network_map/beginning_and_terminus.map",
		"beginning",
		"terminus",
		"20",
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := app.Run(args, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d, stderr=%q", code, stderr.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected empty stderr, got %q", stderr.String())
	}

	out := strings.TrimSpace(stdout.String())
	if out == "" {
		t.Fatalf("expected non-empty movement output")
	}
	turns := strings.Count(out, "\n") + 1
	if turns <= 1 {
		t.Fatalf("expected more than 1 turn, got %d", turns)
	}
	if turns > 11 {
		t.Fatalf("expected at most 11 turns, got %d", turns)
	}
}
