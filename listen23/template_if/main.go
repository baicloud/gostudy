package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Address struct {
	City     string
	Province string
}

type UserInfo struct {
	Name    string
	Sex     string
	Age     int
	Address Address
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Fprintf(w, "load login.html failed, err:%v", err)
		return
	}

	user := UserInfo{
		Name: "Mary",
		Sex:  "male",
		Age:  18,
		Address: Address{
			City:     "beijing",
			Province: "beijing",
		},
	}
	err = t.Execute(w, user)
	if err != nil {
		fmt.Printf("execute template failed,err:%v\n", err)
	}
	t.Execute(os.Stdout, user)
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("listen server failed,err:%v\n", err)
		return
	}
}
