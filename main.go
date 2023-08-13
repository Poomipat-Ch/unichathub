package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	t := template.Must(template.ParseFiles("templates/index.html"))

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}).Methods("GET")

	router.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Button was clicked!")
	}).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
