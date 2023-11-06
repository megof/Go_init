package main

import (
	"fmt" //formato de entrada y salida
	"html/template"
	"net/http" // mostrar una página
	//"log" // mostrar información por consola
)

type User struct {
	Name   string
	Skills string
	Age    int
}

func index(rw http.ResponseWriter, r *http.Request) {
	//template, _ := template.ParseFiles("templates/index.html") //así capturo el error
	template, err := template.ParseFiles("templates/index.html")

	user1 := User{"Meg", "Dev", 2020}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, user1)
	}
}

func main1() {
	http.HandleFunc("/", index)

	fmt.Println("dev: corriendo correctamente en el localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}
