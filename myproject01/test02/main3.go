package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
		_ "github.com/go-sql-driver/mysql"
	"database/sql"
)


func Insert(username,password string){
	//connect database  连接数据库
	db, err1 := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	if err1 != nil{
		fmt.Println(err1)
	}

	//insert data into database   向数据库插入数据
   db.Exec("insert into user2(username,password) values("+username+","+password+")")
}


//处理器函数  第一个参数是一个ResponseWriter接口，第二个参数是一个指向Request的指针
func login(w http.ResponseWriter, r *http.Request){
	 //解析form
	fmt.Println("method: ", r.Method)
	//判断传递方式
	if r.Method == "GET"{
		//加载界面模板
		t, _ := template.ParseFiles("login.html")
		//将解析好的模板应用到data上，这里data为nil
		t.Execute(w, nil)
	}else if r.Method == "POST"{
		//调用函数ParseForm对请求进行语法分析
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"][0])
		fmt.Println("password: ", r.Form["password"][0])
        Insert(r.Form["username"][0],r.Form["password"][0])
	}
}
//ServeHTTP
//func Start( ){
//	//var w http.ResponseWriter
//	//var r *http.Request
//	http.HandleFunc("/", login)
//	err := http.ListenAndServe(":9090", nil)
//	if err != nil{
//		log.Fatal("ListenAndServe: ", err)
//	}
//}

func main(){
	//db, err1 := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	//if err1 != nil{
	//	fmt.Println(err1)
	//}
     //将login函数转换成一个handler，并且将它与DefaultServeMux进行绑定，以此简化创建并绑定Handler的工作
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}

	//Start()
	//var r *http.Request
	//r.ParseForm()
	//fmt.Println(r.Form["username"])
}


