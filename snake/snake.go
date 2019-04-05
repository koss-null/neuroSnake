package snake

import (
	"../utils"
	"../utils/dataStructures"
	"errors"
	"neuroSnake/field"
)

type SnakeMove int8

const (
	None SnakeMove = iota
	Up
	Down
	Left
	Right
)

type SnakeDeadError error

var (
	OutOfBorder SnakeDeadError = errors.New("the Snake got out of border")
	EatsItself  SnakeDeadError = errors.New("the Snake bit it's tail")
)

type (
	snake struct {
		body     dataStructures.LinkedList
		lastMove SnakeMove
		field    field.Field
	}

	Snake interface {
		Move(SnakeMove) SnakeDeadError
		AutoMove() SnakeDeadError
		GetSnake() []utils.Dot2
	}
)

func deriveMove(cur utils.Dot2, move SnakeMove) utils.Dot2 {
	var nextDot utils.Dot2
	switch move {
	case Up:
		nextDot = utils.Dot2{cur.X - 1, cur.Y}
	case Down:
		nextDot = utils.Dot2{cur.X + 1, cur.Y}
	case Left:
		nextDot = utils.Dot2{cur.X, cur.Y - 1}
	case Right:
		nextDot = utils.Dot2{cur.X, cur.Y + 1}
	default:
		nextDot = cur
	}

	return nextDot
}

// returns true if closure exists
func (s *snake) closureCheck() bool {
	exist := map[utils.Dot2]interface{}{}
	for _, dot := range s.GetSnake() {
		if _, ok := exist[dot]; ok {
			return true
		}
		exist[dot] = struct{}{}
	}
	return false
}

func (s *snake) Move(nextMove SnakeMove) SnakeDeadError {
	curDot := s.body.Head().(utils.Dot2)
	nextDot := deriveMove(curDot, nextMove)

	moveRes := s.field.CheckMove(nextDot)

	// we do know that nestDot is a utils.Dot2 type
	_ := s.body.PushFront(nextDot)
	closureExist := s.closureCheck()

	switch {
	case closureExist:
		return EatsItself
	case moveRes != field.GotApple:
		s.body.Pop()
	case moveRes == field.OutOfBorders:
		return OutOfBorder
	}

	return nil
}

func (s *snake) AutoMove() SnakeDeadError {
	return s.Move(s.lastMove)
}

func (s *snake) GetSnake() []utils.Dot2 {
	sl := s.body.Slice()
	dotSl := make([]utils.Dot2, len(sl))
	for i := range sl {
		dotSl[i] = sl[i].(utils.Dot2)
	}
	return dotSl
}
