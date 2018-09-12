package main

import (
	"math/rand"
	"time"
)
var (
     snake Snake
     score int64 = 1
)

func main(){

	//地图初始化
	mapInit()

	//随机数种子
	rand.Seed(time.Now().UnixNano())

	//食物初始化
	var food Food
	food.FoodInit()

	//蛇初始化
	snake.SnakeInit()
	snake.playGame(&food)

	//获取键盘输入
	getInput()



}