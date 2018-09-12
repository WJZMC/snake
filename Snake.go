package main

import (
	"github.com/tncardoso/gocurses"
	"fmt"
	"time"
)

type Snake struct{
	direction int
	body []Point
}

func (s *Snake) SnakeInit() {
	s.direction = 'd'
	s.body = make([]Point, 2)
	s.body[1].x = MaxWith/2.0 - 1
	s.body[1].y = MaxHeight / 2.0
	s.body[0].x = MaxWith / 2.0
	s.body[0].y = MaxHeight / 2.0
	s.refushSnake()
}
func (s *Snake)refushSnake(){
	for i :=0 ; i<len(snake.body);i++{
		if i==0{
			DrawOnly(snake.body[i].x,snake.body[i].y,'@')
			refushWin()
		}else{
			DrawOnly(snake.body[i].x,snake.body[i].y,'#')
			refushWin()
		}
	}
	//refushWin()
}

func (s *Snake)playGame(food *Food){
	offSetX ,offSetY ,lasX,lasY := 0,0,0,0

	go func() {
	FLAG:

		for{
			switch  snake.direction {
			case 'w':
				offSetX = 0
				offSetY = -1

			case 'x':
				offSetX = 0
				offSetY = 1

			case 'a':
				offSetX = -1
				offSetY = 0


			case 'd':
				offSetX = 1
				offSetY = 0
			case 'p':
				goto FLAG
			default:
			}

			time.Sleep(time.Second/time.Duration(score+2))

			lasX=snake.body[len(snake.body)-1].x
			lasY=snake.body[len(snake.body)-1].y

			if snake.body[0].x==food.x&&snake.body[0].y==food.y{
				snake.body=append(snake.body,Point{lasX,lasY})
				food.FoodInit()
				snake.refushSnake()
				score ++
				gocurses.Mvaddstr(1, 9, fmt.Sprintf("%d",score-1))
				gocurses.Refresh()
			}else{
				for i:=len(snake.body)-1;i>=0;i--{
					if i==0{
						snake.body[i].x+=offSetX
						snake.body[i].y+=offSetY
					}else{
						snake.body[i].x=snake.body[i-1].x
						snake.body[i].y=snake.body[i-1].y
					}

				}
				DrawOnly(lasX,lasY,' ')
				snake.refushSnake()
			}

			if snake.body[0].x<1||snake.body[0].x>MaxWith-2 ||snake.body[0].y>MaxHeight-2||snake.body[0].y<1{
				GameOver()
				return
			}
		}
	}()
}