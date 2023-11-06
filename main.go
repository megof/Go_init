package main

import (
	//formato entrada y salida
	"log"           //datos por consola
	"net/http"      //mostrar web
	"text/template" //
)

var templates = template.Must(template.ParseGlob("templates/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola desde index")
	templates.ExecuteTemplate(w, "index.tm", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola desde index")
	templates.ExecuteTemplate(w, "cre.tm", nil)
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/create", Create)

	log.Println("servidor corriendo.")
	http.ListenAndServe(":8080", nil)
}
