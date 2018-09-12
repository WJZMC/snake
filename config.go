package main

import "time"

type Point struct{
	x int
	y int
}

const(
	MaxWith = 80
	MaxHeight = 20
)

func GameMsgAlert(msg string)  {
	Window.Mvaddstr(MaxHeight/2.0, MaxWith/2.0-5, msg)
	Window.Refresh()
}
func GameOver(){
	snake.direction='p'
	GameMsgAlert("Game Over")
	time.Sleep(time.Second*3)
}