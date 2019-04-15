package field

import (
	"fmt"
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
		Apple() utils.Dot2
		Dimensions() (int, int)

		GetMoveChecker() *func(dot utils.Dot2) MoveResult
		GetAppleSetTrigger() *func()
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
	defer f.setApple()
	return &f
}

func (f *field) setApple() {
	// todo: need to try before apple will be not on snake position

	setApple := func() {
		f.apple.X = int(f.randGen.Int31n(int32(f.Width)))
		f.apple.Y = int(f.randGen.Int31n(int32(f.Height)))
	}

	setApple()
}

func (f *field) GetAppleSetTrigger() *func() {
	setAppleFunc := f.setApple
	return &setAppleFunc
}

func (f *field) Apple() utils.Dot2 {
	return f.apple
}

func (f *field) Dimensions() (int, int) {
	return f.Width, f.Height
}

func (f *field) checkMove(dot utils.Dot2) MoveResult {
	fmt.Printf("%d < 0 || %d >= %d || %d < 0 || %d >= %d\n", dot.X, dot.X, f.Width, dot.Y, dot.Y, f.Height)
	switch {
	case dot.X < 0 || dot.X >= f.Width || dot.Y < 0 || dot.Y >= f.Height:
		return OutOfBorders
	case dot.Eq(f.apple):
		return GotApple
	default:
		return OK
	}
}


func (f *field) GetMoveChecker() *func(utils.Dot2) MoveResult {
	moveChecker := f.checkMove
	return &moveChecker
}