package userInput

import (
	"errors"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/eiannone/keyboard"
	"neuroSnake/controller"
	"neuroSnake/field"
	"neuroSnake/snake"
	"neuroSnake/utils"
	"time"
)

type (
	userInput struct {
		speed time.Duration
		Snake *snake.Snake
		Field *field.Field
	}
)

func NewUserInputRunner(speed time.Duration, s *snake.Snake, f *field.Field) controller.Runner {
	return &userInput{speed, s, f}
}

func (r *userInput) Run() chan error {
	errorsBag := make(chan error, 0)

	go func() {
		for {
			rn, _, err := keyboard.GetSingleKey()
			if err != nil {
				errorsBag <- err
			}
			switch rn {
			case 'w':
				(*r.Snake).SetMove(snake.Up)
			case 's':
				(*r.Snake).SetMove(snake.Down)
			case 'a':
				(*r.Snake).SetMove(snake.Left)
			case 'd':
				(*r.Snake).SetMove(snake.Right)
			case 'q':
				errorsBag <- errors.New("game finished")
			}
		}
	}()

	ticker := time.NewTicker(r.speed)
	go func() {
		for {
			select {
			case cash := <-errorsBag:
				errorsBag <- cash
				return
			case <-ticker.C:
				{
					err := (*r.Snake).AutoMove()
					r.Draw()
					if err != nil {
						errorsBag <- err
					}
				}
			}
		}
	}()

	return errorsBag
}

func (r *userInput) Draw() {
	width, height := (*r.Field).Dimensions()
	smap := (*r.Snake).GetSnakeMap()

	tm.Flush()
	defer tm.Clear()

	for i := 0; i <= height; i++ {
		fmt.Printf("%d:\t|", i)
		for j := 0; j <= width; j++ {
			_, isSnake := smap[utils.Dot2{j, i}]
			switch {
			case isSnake:
				fmt.Print("O")
			case j == width:
				fmt.Print("|")
			case i == height:
				fmt.Print("-")
			case (*r.Field).Apple().Eq(utils.Dot2{j, i}):
				fmt.Print("A")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println((*r.Field).Apple())
}
