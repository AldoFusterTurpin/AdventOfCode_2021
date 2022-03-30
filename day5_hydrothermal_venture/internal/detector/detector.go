package detector

type PairCoordinates struct {
	InitCoordinate Coordinate
	EndCoordinate  Coordinate
}

type Coordinate struct {
	X, Y int
}

func GetNumberOfOverlappingPoints(pc []PairCoordinates, overlappingTarget int) int {
	// key: a coordinate
	// value: how many lines intersect at that point
	representation := buildRepresentation(pc)
	return calculateOverlappingPoints(representation, overlappingTarget)
}

func calculateOverlappingPoints(representation map[Coordinate]int, overlappingTarget int) int {
	c := 0

	for _, v := range representation {
		if v >= overlappingTarget {
			c++
		}
	}

	return c
}

func pairIsHorizontalOrVertical(pc PairCoordinates) bool {
	return pairIsHorizontal(pc) || pairIsVertical(pc)
}

func pairIsHorizontal(pc PairCoordinates) bool {
	return pc.InitCoordinate.X == pc.EndCoordinate.X
}

func pairIsVertical(pc PairCoordinates) bool {
	return pc.InitCoordinate.Y == pc.EndCoordinate.Y
}

func buildRepresentation(pc []PairCoordinates) map[Coordinate]int {
	representation := make(map[Coordinate]int)

	for _, pair := range pc {
		if pairIsHorizontalOrVertical(pair) {
			treatPair(pair, representation)
		}
	}

	return representation
}

func treatPair(p PairCoordinates, representation map[Coordinate]int) {
	if pairIsHorizontal(p) {
		buildHorizontalLine(p.InitCoordinate, p.EndCoordinate, representation)
	} else {
		buildVerticalLine(p.InitCoordinate, p.EndCoordinate, representation)
	}
}

func buildVerticalLine(initCoordinate Coordinate, endCoordinate Coordinate, representation map[Coordinate]int) {
	start, end := getStartAndEnd(initCoordinate.X, endCoordinate.X)

	for i := start; i <= end; i++ {
		p := Coordinate{
			X: i,
			Y: initCoordinate.Y,
		}
		representation[p]++
	}
}

func buildHorizontalLine(initCoordinate Coordinate, endCoordinate Coordinate, representation map[Coordinate]int) {
	start, end := getStartAndEnd(initCoordinate.Y, endCoordinate.Y)

	for i := start; i <= end; i++ {
		p := Coordinate{
			X: initCoordinate.X,
			Y: i,
		}
		representation[p]++
	}
}

func getStartAndEnd(x1, x2 int) (start int, end int) {
	if x1 < x2 {
		return x1, x2
	}
	return x2, x1
}
