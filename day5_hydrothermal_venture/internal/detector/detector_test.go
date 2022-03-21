package detector_test

import (
	"hydrothermalventure/internal/detector"
	"reflect"
	"testing"
)

func TestBuildVentsDiagram(t *testing.T) {
	tests := map[string]struct {
		pc              []detector.PairCoordinates
		expectedDiagram []string
	}{
		"statement_sampe": {
			pc: []detector.PairCoordinates{
				detector.PairCoordinates{
					detector.Coordinate{X: 0, Y: 9},
					detector.Coordinate{X: 5, Y: 9},
				},
				detector.PairCoordinates{
					detector.Coordinate{X: 8, Y: 0},
					detector.Coordinate{X: 0, Y: 8}},
				detector.PairCoordinates{
					detector.Coordinate{X: 9, Y: 4},
					detector.Coordinate{X: 3, Y: 4}},
				detector.PairCoordinates{
					detector.Coordinate{X: 2, Y: 2},
					detector.Coordinate{X: 2, Y: 1}},
				detector.PairCoordinates{
					detector.Coordinate{X: 7, Y: 0},
					detector.Coordinate{X: 7, Y: 4}},
				detector.PairCoordinates{
					detector.Coordinate{X: 6, Y: 4},
					detector.Coordinate{X: 2, Y: 0}},
				detector.PairCoordinates{
					detector.Coordinate{X: 0, Y: 9},
					detector.Coordinate{X: 2, Y: 9}},
				detector.PairCoordinates{
					detector.Coordinate{X: 3, Y: 4},
					detector.Coordinate{X: 1, Y: 4}},
				detector.PairCoordinates{
					detector.Coordinate{X: 0, Y: 0},
					detector.Coordinate{X: 8, Y: 8}},
				detector.PairCoordinates{
					detector.Coordinate{X: 5, Y: 5},
					detector.Coordinate{X: 8, Y: 2}},
			},
			expectedDiagram: []string{
				".......1..",
				"..1....1..",
				"..1....1..",
				".......1..",
				".112111211",
				"..........",
				"..........",
				"..........",
				"..........",
				"222111....",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := detector.BuildRepresentation(tc.pc)

			if !reflect.DeepEqual(got, tc.expectedDiagram) {
				t.Errorf("%s: expected diagram %v, but got: %v", name, tc.expectedDiagram, got)
			}
		})
	}
}
