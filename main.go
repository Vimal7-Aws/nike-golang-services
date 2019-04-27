package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/rivers", riverHandler).Methods("GET")
	r.HandleFunc("/mountains", mountainHandler).Methods("GET")

	fmt.Println("Starting the server at 2050")

	http.ListenAndServe(":2050", r)
}

func homeHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Print("home handler")
}

func riverHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Print("home handler")
}

func mountainHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Print("home handler")
}