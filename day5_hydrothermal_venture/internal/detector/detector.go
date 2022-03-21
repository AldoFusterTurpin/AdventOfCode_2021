package detector

type PairCoordinates struct {
	InitCoordinate Coordinate
	EndCoordinate  Coordinate
}

type Coordinate struct {
	X, Y int
}

func BuildRepresentation(pc []PairCoordinates) []string {
	return []string{
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
	}
}
