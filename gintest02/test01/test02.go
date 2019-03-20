package main

type Test2 struct{
	Id int64
	Name string
	Version int `xorm:"version"`
}

//const base_format = "2006-01-02 15:04:05"

func main(){

}

//func StringTimeToDateTime(stime string)time.Time{
//	loc, _ := time.LoadLocation("Asia/Shanghai")
//	tTime, _ := time.ParseInLocation(base_format, stime, loc)
//	return tTime
//}
//
//func VersionTest(){
//
//}