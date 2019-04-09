package snake

import (
	"errors"
	"neuroSnake/utils"
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

type Snake interface {
		Move(SnakeMove) SnakeDeadError
		AutoMove() SnakeDeadError
		SetMove(SnakeMove)
		GetSnake() []utils.Dot2
		GetSnakeMap() map[utils.Dot2]interface{}
	}
