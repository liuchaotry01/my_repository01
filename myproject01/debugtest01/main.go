package main

import "fmt"

func main(){
	var x, y int
	y = 0
	for i :=0;i < 10; i++{
		x = i
		y +=x
		fmt.Println(x,y)
	}
}
