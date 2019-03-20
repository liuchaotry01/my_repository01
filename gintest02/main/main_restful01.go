package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func main(){
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
}
