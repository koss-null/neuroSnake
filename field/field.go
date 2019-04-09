package field

import (
	"math/rand"
	"neuroSnake/utils"
	"time"
)

type (
	field struct {
		Width   int
		Height  int
		apple   utils.Dot2
		randGen *rand.Rand
	}

	Field interface {
		CheckMove(dot utils.Dot2) MoveResult
		SetApple()
		Apple() utils.Dot2
		Dimensions() (int, int)
	}
)

type MoveResult int

const (
	OK MoveResult = iota
	OutOfBorders
	GotApple
)

func MakeField(w int, h int) Field {
	if w <= 0 || h <= 0 {
		// fixme
		panic("invalid field size")
	}

	f := field{
		Width:   w,
		Height:  h,
		randGen: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	defer f.SetApple()
	return &f
}


func (f *field) SetApple() {
	setApple := func() {
		f.apple.X = int(f.randGen.Int31n(int32(f.Width)))
		f.apple.Y = int(f.randGen.Int31n(int32(f.Height)))
	}

	setApple()
}

func (f *field) Apple() utils.Dot2 {
	return f.apple
}

func (f *field) Dimensions() (int, int) {
	return f.Width, f.Height
}

func (f *field) CheckMove(dot utils.Dot2) MoveResult {
	switch {
	case dot.X < 0 || dot.X > f.Width || dot.Y < 0 || dot.Y > f.Height:
		return OutOfBorders
	case dot.Eq(f.apple):
		return GotApple
	default:
		return OK
	}
}
