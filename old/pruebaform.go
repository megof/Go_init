package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Análizar los parámetros URL enviados, entonces analizar el paquete para analizar el cuerpo del paquete, para peticiones POST.
	// precaución: Si no se llama al método ParseForm, la siguiente unformación no podra ser obtenida del Formulario
	fmt.Println(r.Form) // Imprime la información del lado del servidor
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // Escrite la respuesta
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// Parte lógica de la función login
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main2() {
	http.HandleFunc("/", sayhelloName) // Definimos la regla del enrutador
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // Definimos el puerto de escucha
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
