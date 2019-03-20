package main

import "fmt"

type Test1 struct {
	SpreadId string
	UserId string
	IsDel int
}

func main() {
	var a,b float32
	a = 0.000000000000000000000000000000000000000000000000000000000000000000001
	b = 0.0
	//if a > 0{
	//	fmt.Println(a,b)
	//}
	//if b < 0{
	//	fmt.Println(b)
	//}
	//if a >= 0{
	//	fmt.Println(a)
	//}
	if a == b{
		fmt.Println("01",a)
	}
}
