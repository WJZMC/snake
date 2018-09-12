package main

import (
	"github.com/tncardoso/gocurses"
	"fmt"
)

var(
	Window *gocurses.Window
)
func DrawOnly(x,y int,ch int){
	Window.Mvaddstr(y, x, fmt.Sprintf("%c",ch))
}
func refushWin(){
	Window.Refresh()
}
func mapInit(){
	fmt.Println("starting")
	gocurses.Initscr()
	defer gocurses.End()
	gocurses.Cbreak()
	gocurses.Noecho()
	gocurses.CursSet(0)
	gocurses.Stdscr.Keypad(true)

	y, x := gocurses.Getmaxyx()
	//gocurses.Addstr("Press any key to exit",x,y)
	gocurses.Mvaddstr(0, MaxWith/2.0-2, "snake")
	gocurses.Mvaddstr(1, 1, "score:  0")
	gocurses.Mvaddstr(1, 80-37, "a:left w:up x:down d:right esc:exit")
	gocurses.Refresh()


	Window = gocurses.NewWindow(MaxHeight, MaxWith, (y-MaxHeight)/2, (x-MaxWith)/2)
	Window.Box(0, 0)
	Window.Refresh()
}
func getInput(){
	for {
		ch := gocurses.Getch()

		switch  ch {
		case 'w':
			if snake.direction == 'x'{
				GameOver()
				return
			}else{
				snake.direction = 'w'

			}

		case 'x':
			if snake.direction == 'w'{
				GameOver()
				return
			}else {
				snake.direction = 'x'
			}

		case 'a':
			if snake.direction == 'd'{
				GameOver()
				return
			}else {
				snake.direction = 'a'
			}

		case 'd':
			if snake.direction == 'a'{
				GameOver()
				return
			}else {
				snake.direction = 'd'
			}
		case 'p':
			snake.direction='p'
		case 27:
			return
		default:
		}
	}
}
