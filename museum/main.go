package main

import (
	"fmt"
	"html/template"
	"net/http"

	"example.com/museum/api"
	"example.com/museum/data"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from new go server!"))
}

func anotherHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Another page with hello"))
}

// go template
// parse content into html
func handleTemplate(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}
	// this string will be embbeded in html
	template.Execute(w, data.GetAll())
}

// http server
func main() {
	// multiplexer
	mux := http.NewServeMux()
	// serves HTTP requests with the contents of the file system rooted at root.
	frontHandler := http.FileServer(http.Dir("./public"))
	mux.Handle("/", frontHandler)
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/hello/another", anotherHello)
	mux.HandleFunc("/template", handleTemplate)
	mux.HandleFunc("GET /api/exhibitions", api.Get)
	mux.HandleFunc("GET /api/exhibitions/{id}", api.Get)
	mux.HandleFunc("POST /api/exhibitions", api.Post)

	err := http.ListenAndServe(":3333", mux)

	if err != nil {
		fmt.Printf("something happened with server: %v", err.Error())
	}

}
