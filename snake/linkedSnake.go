package snake

import (
	"neuroSnake/field"
	"neuroSnake/utils"
	"neuroSnake/utils/dataStructures"
)

type snake struct {
	body            *dataStructures.LinkedList
	lastMove        SnakeMove
	checkMove       *func(utils.Dot2) field.MoveResult
	triggerAppleSet *func()
}

func deriveMove(cur utils.Dot2, move SnakeMove) utils.Dot2 {
	var nextDot utils.Dot2
	switch move {
	case Up:
		nextDot = utils.Dot2{cur.X, cur.Y - 1}
	case Down:
		nextDot = utils.Dot2{cur.X, cur.Y + 1}
	case Left:
		nextDot = utils.Dot2{cur.X - 1, cur.Y}
	case Right:
		nextDot = utils.Dot2{cur.X + 1, cur.Y}
	default:
		nextDot = cur
	}

	return nextDot
}

/*
	snake need to be provided with two functions:
	checkMove function which returns some move result depends of field parameters
	appleSer function which triggers change of an apple position
 */
func MakeSnake(initDot utils.Dot2, cm *func(utils.Dot2) field.MoveResult, as *func()) Snake {
	body := dataStructures.NewLinkedList()
	err := body.PushBack(initDot)
	if err != nil {
		panic(err)
	}

	// fixme: Snake must start moving to a random side, not to the right
	return &snake{
		&body,
		Right,
		cm,
		as,
	}
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
	curDot := (*s.body).Head().(utils.Dot2)
	nextDot := deriveMove(curDot, nextMove)

	moveRes := (*s.checkMove)(nextDot)

	// we do know that nestDot is a utils.Dot2 type
	_ = (*s.body).PushFront(nextDot)
	closureExist := s.closureCheck()

	switch {
	case closureExist:
		return EatsItself
	case moveRes == field.OutOfBorders:
		return OutOfBorder
	case moveRes == field.GotApple:
		(*s.triggerAppleSet)()
	default:
		(*s.body).Pop()
	}

	return nil
}

func (s *snake) AutoMove() SnakeDeadError {
	return s.Move(s.lastMove)
}

func (s *snake) SetMove(sm SnakeMove) {
	s.lastMove = sm
}

func (s *snake) GetSnake() []utils.Dot2 {
	sl := (*s.body).Slice()
	dotSl := make([]utils.Dot2, len(sl))
	for i := range sl {
		dotSl[i] = sl[i].(utils.Dot2)
	}
	return dotSl
}

func (s *snake) GetSnakeMap() map[utils.Dot2]interface{} {
	smap := map[utils.Dot2]interface{}{}
	for _, dot := range s.GetSnake() {
		smap[dot] = struct{}{}
	}
	return smap
}
