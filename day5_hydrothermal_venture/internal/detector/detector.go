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

// func iterateBetweenTwoCoordinates(c1 Coordinate, c Coordinate, representation map[Coordinate]int) {
// 	for i := p.InitCoordinate.X; i < p.EndCoordinate.X; i++ {
// 		p := Coordinate{
// 			X: i,
// 			Y: p.InitCoordinate.Y,
// 		}
// 		representation[p]++
// 	}
// }

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
		start, end := getStartAndEnd(p.InitCoordinate.Y, p.EndCoordinate.Y)

		for i := start; i <= end; i++ {
			p := Coordinate{
				X: p.InitCoordinate.X,
				Y: i,
			}
			representation[p]++
		}
	} else {
		start, end := getStartAndEnd(p.InitCoordinate.X, p.EndCoordinate.X)

		for i := start; i <= end; i++ {
			p := Coordinate{
				X: i,
				Y: p.InitCoordinate.Y,
			}
			representation[p]++
		}
	}
}

func getStartAndEnd(x1, x2 int) (start int, end int) {
	if x1 < x2 {
		return x1, x2
	}
	return x2, x1
}
