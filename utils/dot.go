package utils

type (
	Dot2 struct {
		X, Y int
	}
)

func (dot Dot2) Eq(otherDot Dot2) bool {
	return dot.x == otherDot.x && dot.y == otherDot.y
}

