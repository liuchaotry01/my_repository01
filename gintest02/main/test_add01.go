package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"strconv"
)

//type Database struct{
//	DB *sql.DB
//}

var Db *sql.DB

type Person struct{
	Id         int     `json: "id" form: "id"`
	FirstName  string  `json:"first_name" form:"first_name"`
	LastName   string  `json:"last_name" form:"last_name"`
}

func main(){
	//db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mysql?parseTime=true")
	//Db = db
	//if err != nil{
	//	log.Fatalln(err)
	//}
	//defer db.Close()
	//
	//
	//db.SetMaxIdleConns(20)
	//db.SetMaxOpenConns(20)
	//
	//if err := db.Ping(); err != nil{
	//	log.Fatalln(err)
	//}

	InitDb()
	defer Db.Close()
	//db := GetDb()

	router := gin.Default()

	router.POST("/person", Add)
	router.GET("/persons",MyQuery)
	router.GET("/person/:id", MyQueryOneRecord)
	router.PUT("liu/update/:id",MyUpdate)
	router.DELETE("/delete/:id",MyDelete)

	router.Run(":8090")
}

func InitDb(){
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mysql?parseTime=true")
	if err != nil{
		log.Fatalln(err)
	}
	//defer db.Close()


	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil{
		log.Fatalln(err)
	}

	Db = db
	//for i := 0; i<2; i++{}
}

func GetDb()*sql.DB{
	return Db
}

func Add(c *gin.Context){
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	db :=GetDb()
	rs, err := db.Exec("INSERT INTO person(first_name, last_name)VALUES(?,?)",firstName,lastName)
	if err != nil{
		log.Fatalln(err)
	}

	id, err := rs.LastInsertId()
	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println("insert person Id{}", id)
	msg := fmt.Sprintf("insert successful %d", id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func MyQuery(c *gin.Context){
	db := GetDb()
	rows, err := db.Query("SELECT id,first_name,last_name FROM person")

	if err != nil{
		log.Fatalln(err)
	}
	defer rows.Close()

	persons :=make([]Person,0)
	for rows.Next(){
		var person Person
		rows.Scan(&person.Id,&person.FirstName,&person.LastName)
		persons = append(persons,person)
	}

	if err = rows.Err();err != nil{
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK,gin.H{
		"persons": persons,
	})
}

func MyQueryOneRecord(c *gin.Context){
	db := GetDb()
	id := c.Param("id")
	var person Person
	err := db.QueryRow("SELECT id, first_name, last_name FROM person WHERE id=?", id).Scan(
		&person.Id, &person.FirstName, &person.LastName,)

	if err != nil{
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"person": nil,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"person": person,
	})
}

func MyUpdate(c *gin.Context){
	db := GetDb()
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	person := Person{Id: id}
	err = c.Bind(&person)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(person)
	stmt, err := db.Prepare("UPDATE person SET first_name=?, last_name=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		log.Fatalln(err)
	}
	rs, err := stmt.Exec(person.FirstName, person.LastName, person.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Update person %d successful %d", person.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func MyDelete(c *gin.Context){
	db := GetDb()
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil{
		log.Fatalln(err)
	}

	rs, err := db.Exec("DELETE FROM person WHERE id=?",id)
	if err != nil{
		log.Fatalln(err)
	}

	ra, err := rs.RowsAffected()
	if err != nil{
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Delete person %d successful %d", id, ra)
	c.JSON(http.StatusOK,gin.H{
		"msg": msg,
	})
}