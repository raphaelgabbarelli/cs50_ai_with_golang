package search

type Point struct {
	X, Y int
}

func (p Point) Equal(b any) bool {
	if bPoint, ok := b.(Point); ok {
		return p.X == bPoint.X && p.Y == bPoint.Y
	}
	return false // consider throwing some error
}
