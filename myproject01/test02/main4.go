package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err1 := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	if err1 != nil{
		fmt.Println(err1)
	}
	ret, _ :=db.Exec("insert into user2(id) values(1)")
	fmt.Println(ret)
}
