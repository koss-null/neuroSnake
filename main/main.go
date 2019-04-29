package main

import (
	"fmt"
	"neuroSnake/controller/userInput"
	"neuroSnake/field"
	"neuroSnake/snake"
	"neuroSnake/utils"
	"time"
)

func main() {
	fld := field.MakeField(20, 10)
	snk := snake.MakeSnake(utils.Dot2{10, 5}, fld.GetMoveChecker(), fld.GetAppleSetTrigger())
	//fld.SetSnake(&snk)
	exitChan := userInput.NewUserInputRunner(time.Millisecond*300, &snk, &fld).Run()
	fmt.Println(<-exitChan)
}
