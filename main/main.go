package main

import (
	"neuroSnake/field"
	"neuroSnake/game"
	"neuroSnake/snake"
	"time"
)

func main() {
	fld := field.MakeField(20, 20)
	snk := snake.MakeSnake(&fld)
	//fld.SetSnake(&snk)
	exitChan := game.NewRunner(time.Millisecond * 500, &snk, &fld).Run()
	for {
		<-exitChan
	}
}
