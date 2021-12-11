package dive_test

import (
	"dive_mod/dive"
	"testing"
)

func TestProcessCommands(t *testing.T) {
	tests := map[string]struct {
		input             []string
		expectedHPosition int
		expectedDepth     int
	}{
		"simple": {
			input:             []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			expectedHPosition: 15,
			expectedDepth:     10,
		},
	}

	for name, tc := range tests {
		hPosition, depth := dive.ProcessCommands(tc.input)
		if hPosition != tc.expectedHPosition {
			t.Fatalf("%s: expected: %v, but got: %v", name, tc.expectedHPosition, hPosition)
		}
		if depth != tc.expectedDepth {
			t.Fatalf("%s: expected: %v, but got: %v", name, tc.expectedDepth, depth)
		}
	}
}

// Part 2
func TestProcessCommandsWithAim(t *testing.T) {
	tests := map[string]struct {
		input             []string
		expectedHPosition int
		expectedDepth     int
	}{
		"simple": {
			input:             []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			expectedHPosition: 15,
			expectedDepth:     60,
		},
	}

	for name, tc := range tests {
		hPosition, depth := dive.ProcessCommandsWithAim(tc.input)
		if hPosition != tc.expectedHPosition {
			t.Fatalf("%s: expected: %v, but got: %v", name, tc.expectedHPosition, hPosition)
		}
		if depth != tc.expectedDepth {
			t.Fatalf("%s: expected: %v, but got: %v", name, tc.expectedDepth, depth)
		}
	}
}
