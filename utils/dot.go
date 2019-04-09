package utils

type (
	Dot2 struct {
		X, Y int
	}
)

func (dot Dot2) Eq(otherDot Dot2) bool {
	return dot.X == otherDot.X && dot.Y == otherDot.Y
}

