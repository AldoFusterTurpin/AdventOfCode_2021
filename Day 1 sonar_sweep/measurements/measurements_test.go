package measurements_test

import (
	"sonar_sweep/measurements"
	"testing"
)

func TestMeasurementsLargerThanPrevious(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"simple": {
			input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			want:  7,
		},
		"simple2": {
			input: []int{},
			want:  0,
		},
		"simple3": {
			input: []int{1},
			want:  0,
		},
		"simple4": {
			input: []int{199, 200, 300, 302, 1, 2, 3},
			want:  5,
		},
	}

	for name, tc := range tests {
		got := measurements.MeasurementsLargerThanPrevious(tc.input)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, but got: %v", name, tc.want, got)
		}
	}
}

// Part 2
func TestWindowsMeasurementsLargerThanPrevious(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"simple": {
			input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			want:  5,
		},
	}

	for name, tc := range tests {
		got := measurements.WindowsMeasurementsLargerThanPrevious(tc.input)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, but got: %v", name, tc.want, got)
		}
	}
}
