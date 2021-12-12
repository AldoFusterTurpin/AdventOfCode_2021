package diagnostic_test

import (
	"binary_diagnostic/diagnostic"
	"testing"
)

func TestProcessBinaryNumbers(t *testing.T) {
	tests := map[string]struct {
		input               []string
		expectedGammaRate   int64
		expectedEpsilonRate int64
	}{
		"simple": {
			input:               []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			expectedGammaRate:   22,
			expectedEpsilonRate: 9,
		},
	}

	for name, tc := range tests {
		gammaRate, epsilonRate := diagnostic.ProcessBinaryNumbers(tc.input)
		if gammaRate != tc.expectedGammaRate {
			t.Fatalf("%s: expected: %v, but got: %v", name, tc.expectedGammaRate, gammaRate)
		}
		if epsilonRate != tc.expectedEpsilonRate {
			t.Fatalf("%s: expected: %v, but got: %v", name, tc.expectedEpsilonRate, epsilonRate)
		}
	}
}

func TestGetOxygenGeneratorRatingAndCO2ScrubberRating(t *testing.T) {
	tests := map[string]struct {
		input                         []string
		expectedOxygenGeneratorRating int64
		expectedCO2ScrubberRating     int64
	}{
		"simple": {
			input:                         []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			expectedOxygenGeneratorRating: 23,
			expectedCO2ScrubberRating:     10,
		},
		"simple2": {
			input:                         []string{"00100", "11110"},
			expectedOxygenGeneratorRating: 30,
			expectedCO2ScrubberRating:     4,
		},
	}

	for name, tc := range tests {
		oxygenGeneratorRating, CO2ScrubberRating := diagnostic.GetOxygenGeneratorRatingAndCO2ScrubberRating(tc.input)
		if oxygenGeneratorRating != tc.expectedOxygenGeneratorRating {
			t.Fatalf("%s: expected: %v, but got: %v", name, tc.expectedOxygenGeneratorRating, oxygenGeneratorRating)
		}
		if CO2ScrubberRating != tc.expectedCO2ScrubberRating {
			t.Fatalf("%s: expected: %v, but got: %v", name, tc.expectedCO2ScrubberRating, CO2ScrubberRating)
		}
	}
}
