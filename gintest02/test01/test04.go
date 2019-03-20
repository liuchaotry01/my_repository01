package main

import (
	"time"
	"fmt"
)

const base_format = "2006-01-02 15:04:05.000000"

func main(){
	endtime := "2018-12-12 12:12:12"
	time,err := StringTimeToDateTime(endtime)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(time)
}

func StringTimeToDateTime(stime string) (time.Time,error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil{
		return time.Now(),err
	}
	tTime, err := time.ParseInLocation(base_format, stime, loc)
	if err != nil{
		return time.Now(),err
	}
	return tTime,nil
}