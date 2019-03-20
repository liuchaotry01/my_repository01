package main

import (
			"net/http"
	"log"
	"fmt"
)

//func sayHelloName(w http.ResponseWriter, r *http.Request){
//	r.ParseForm()
//	fmt.Println(r.Form)
//	fmt.Println("path: ", r.URL.Path)
//	fmt.Println("scheme: ", r.URL.Scheme)
//	fmt.Println(r.Form["url_long"])
//	for k, v := range r.Form{
//		fmt.Println("key: ", k)
//		fmt.Println("val: ", strings.Join(v, " "))
//	}
//	fmt.Fprintf(w, "hello chain!")
//}

//func login(w http.ResponseWriter,r *http.Request)

func Start(){
	//http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}

func main(){
	var w http.ResponseWriter
	var r *http.Request
	Start()
	r.ParseForm()
	fmt.Println(r.Form["username"])
	fmt.Fprintf(w,"hello")
}
