package main

import (
	"neuroSnake/field"
	"neuroSnake/game"
	"neuroSnake/snake"
	"neuroSnake/utils"
	"time"
)

func main() {
	fld := field.MakeField(20, 10)
	snk := snake.MakeSnake(utils.Dot2{10, 5}, fld.GetMoveChecker(), fld.GetAppleSetTrigger())
	//fld.SetSnake(&snk)
	exitChan := game.NewRunner(time.Millisecond * 300, &snk, &fld).Run()
	for {
		<-exitChan
	}
}
