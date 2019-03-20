package main

import (
	"time"
	"fmt"
)

func main(){
		//test()
		TimerDemo := time.NewTimer(time.Second*5)
		for{
			select{
			case <- TimerDemo.C:
				fmt.Println("now time %v",time.Now())
			    TimerDemo.Reset(time.Second*5)
			}
		}
}

func StringTimeToDateTime(stime string) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Now(), err
	}
	tTime, err := time.ParseInLocation(base_format, stime, loc)
	if err != nil {
		return time.Now(), err
	}
	return tTime, nil
}
