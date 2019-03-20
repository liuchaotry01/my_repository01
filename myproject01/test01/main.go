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

	//ret, _ :=db.Exec("insert into user1(name) values('liuchao') where id=?", 1)
	//ins_id, _ := ret.LastInsertId();
	//fmt.Println(ins_id);
	//fmt.Println(ret)

	//ret1, _ :=db.Exec("update user1 set name = 'liu' where id = 1")
    //fmt.Println(ret1)

    //ret2, _ := db.Exec("insert into user1(name) values('liuchao')")
    //fmt.Println(ret2)

    //var id int
	//var name string
    //rows:= db.QueryRow("select id,name from user1 where id = 1")
    //rows.Scan(&id,&name)
    //fmt.Println(id,name)

    //ret3, _ := db.Exec("delete from user1 where id = 2")
    //fmt.Println(ret3)
}
