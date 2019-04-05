package field

import (
	"math/rand"
	"neuroSnake/snake"
	"neuroSnake/utils"
	"time"
)

type (
	field struct {
		Width   int
		Height  int
		apple   utils.Dot2
		snk     *snake.Snake
		randGen *rand.Rand
	}

	Field interface {
		CheckMove(dot utils.Dot2) MoveResult
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

	return &field{
		Width:   w,
		Height:  h,
		randGen: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (f *field) SetSnake(s *snake.Snake) {
	f.snk = s
}

func (f *field) SetApple() {
	exist := map[utils.Dot2]interface{}{}
	for _, dot := range f.snk.GetSnake() {
		exist[dot] = struct{}{}
	}

	setApple := func() {
		f.randGen.Seed(int64(f.Width))
		f.apple.X = int(f.randGen.Uint32())

		f.randGen.Seed(int64(f.Height))
		f.apple.Y = int(f.randGen.Uint32())
	}

	setApple()
	for _, ok := exist[f.apple]; ok; {
		setApple()
	}
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
