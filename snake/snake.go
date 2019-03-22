package snake

import (
	"../utils"
	"../utils/dataStructures"
	"errors"
	"google.golang.org/genproto/googleapis/appengine/v1"
)

type SnakeMove int8

const (
	None  SnakeMove = iota
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
	}

	Snake interface {
		Move(SnakeMove) SnakeDeadError
		Automove() SnakeDeadError
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

func (s *snake) Move(nextMove SnakeMove) SnakeDeadError {
	curDot := s.body.Head().(utils.Dot2)
	nextDot := deriveMove(curDot, nextMove)

	s.body.PushFront(nextDot)
}

func (s *snake) Automove() SnakeDeadError {
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
