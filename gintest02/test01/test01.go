package main

import (
	"github.com/go-xorm/xorm"
	"log"
	_"github.com/go-sql-driver/mysql"
		"time"
	"fmt"
)

type Group1 struct{
	Id int64
	Name string
}

type Test1 struct{
	TestId int64     `xorm:"test_id autoincr"`
	Name string
	GroupId int64   //`xorm:"index"`
	CreateTime  time.Time `xorm:"create_time created" json:"create_time"`
	EndTime   time.Time `xorm:"end_time updated" json:"end_time"`
}

type T1 struct {
	Id   int64
	Name  string
}
//
//type Group1Id struct {
//	Name string
//}


var Engine *xorm.Engine

func init(){
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:liuchao@tcp(localhost:3306)/lcctest01?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := Engine.Sync(new(Test1),new(Group1)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
}

func (t *Test1) TableName() string{
	return "test1"
}

func main(){

     t1 := Test1{
     	Name:"chaohahahahahaha",
     }
	session := Engine.NewSession()
	session.Begin()
	session.InsertOne(&t1)
	fmt.Println(t1)
	session.Commit()
}


//func Transform(baseData []map[string][]byte,typeStruct interface{}) []interface{}{
//	counters:=[]interface{}{}
//	for _,value:=range baseData{
//		dataV:=reflect.Indirect(reflect.ValueOf(typeStruct))
//		count:=dataV.NumField()
//		for k,v:=range value {
//			for i := 0; i < count; i++ {
//				nam := strings.ToLower(dataV.Type().Field(i).Tag.Get(""))
//				if k == nam {
//					typ := dataV.Type().Field(i).Type.Kind()
//					switch typ {
//					case reflect.Int64:{
//						vInt64, _ := strconv.ParseInt(string(v), 10, 64)
//						dataV.Field(i).SetInt(vInt64)
//						break
//					}
//					case reflect.String:{
//						vString := string(v)
//						dataV.Field(i).SetString(vString)
//						break
//					}
//					case reflect.Bool:{
//						vBool, _ := strconv.ParseBool(string(v))
//						dataV.Field(i).SetBool(vBool)
//						break
//					}
//					default:
//						break
//					}
//				}
//			}
//		}
//		//fmt.Println(dataV)
//		dataVV:=dataV.Interface()
//		counters=append(counters,dataVV)
//	}
//	return counters
//}


//func MySelect(){

	//total, _ := Engine.Where("id > ?", 1).Count()
	//fmt.Println(total)
	//ints := MyTest()
	//fmt.Println(ints)

	//gid := make([]Group1Id,0)
	//sql := "select id from group1"
	//res,err := Engine.Exec(sql)
	//var name []string
	//name := make([]string,0)
	//ints := make([]Group1,0)
	//ints1 := make([]Group1,0)
	//err := Engine.Table("group1").Cols("name").Find(&name)
	//fmt.Println(err)
	//fmt.Println(name)
	//err := Engine.SQL("select id,name from group1").Find(&ints)
	//ints1 = ints
	//if err != nil{
	//	fmt.Println(err)
	//}else{
	//	fmt.Println(ints)
	//	fmt.Println(ints1)
	//	fmt.Println(len(ints))
	//}
//}

//func MyTest()(rels []Group1 ){
//	ints := make([]Group1,0)
//	//ints1 := make([]Group1,0)
//	//err := Engine.Table("group1").Cols("id").Find(&ints)
//	err := Engine.SQL("select id,name from group1").Find(&ints)
//	//ints1 = ints
//	if err != nil{
//		fmt.Println(err)
//	}else{
//		fmt.Println("yse")
//		//fmt.Println(ints1)
//		fmt.Println(len(ints))
//	}
//	return ints
//}


//func MyDelete(){
//	sql := "delete from test3 where id = ?"
//	res, err := Engine.Exec(sql,0)
//	if err != nil{
//		fmt.Println(err)
//	}else{
//		fmt.Println(res)
//	}
//}

//func VersionTest(){
//     var test Test3
//     //test.Name = "chao02"
//     Engine.ID(1).Get(&test)
//     rel, err :=Engine.Id(1).Update(&test)
//     //rel, err := Engine.Insert(test)
//     if err != nil{
//     	fmt.Println(err)
//	 }else{
//	 	fmt.Println(rel)
//	 }
//}

//func Join(){
//	tests := make([]TestGroup1,0)
//	err := Engine.Join("INNER","group1","group1.id = test1.group_id").Find(&tests)
//	if err != nil{
//		fmt.Println(err)
//	}else{
//		fmt.Println(tests)
//	}
//}

//func (TestGroup1) TableName() string {
//	return "test1"
//}
