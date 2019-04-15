package game

import (
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/eiannone/keyboard"
	"neuroSnake/field"
	"neuroSnake/snake"
	"neuroSnake/utils"
	"time"
)

type (
	runner struct {
		speed time.Duration
		Snake *snake.Snake
		Field *field.Field
	}

	Runner interface {
		Run() chan interface{}
	}
)

func NewRunner(speed time.Duration, s *snake.Snake, f *field.Field) Runner {
	return &runner{speed, s, f}
}

// todo: make speed depend on level
func (r *runner) Run() chan interface{} {
	closer := make(chan interface{}, 0)

	go func() {
		for {
			rn, _, err := keyboard.GetSingleKey()
			if err != nil {
				// fixme
				// ignore
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
				panic("GG")
			}
		}
	}()

	ticker := time.NewTicker(r.speed)
	go func() {
		for {
			select {
			case <-closer:
				return
			case <-ticker.C:
				{
					err := (*r.Snake).AutoMove()
					r.Draw()
					if err != nil {
						panic(err)
					}
				}
			}
		}	}()

	return closer
}

func (r *runner) Draw() {
	width, height := (*r.Field).Dimensions()
	smap := (*r.Snake).GetSnakeMap()

	tm.Flush()
	defer tm.Clear()

	for i := 0; i <= height; i++ {
		fmt.Printf("%d:\t|", i)
		for j := 0; j <= width; j++ {
			_, isSnake := smap[utils.Dot2{j,	i}]
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