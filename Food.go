package main

import "math/rand"

type Food struct{
	Point
}

func (f *Food)FoodInit(){
	f.x=rand.Intn(MaxWith-2)+1
	f.y=rand.Intn(MaxHeight-2)+1
	DrawOnly(f.x,f.y,'#')
	refushWin()
}