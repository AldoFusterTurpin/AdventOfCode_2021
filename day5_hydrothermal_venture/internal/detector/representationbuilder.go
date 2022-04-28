package detector

type RepresentationBuilder interface {
	buildRepresentation(pc []PairCoordinates) map[Coordinate]int
}

type HorizontalAndVerticalLinesBuilder struct {
}

func (s HorizontalAndVerticalLinesBuilder) buildRepresentation(pc []PairCoordinates) map[Coordinate]int {
	representation := make(map[Coordinate]int)

	for _, pair := range pc {
		if pairIsHorizontalOrVertical(pair) {
			treatPair(pair, representation)
		}
	}

	return representation
}
