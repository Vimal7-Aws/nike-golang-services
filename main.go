package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", healthHandler).Methods("GET")
	r.HandleFunc("/rivers", riverHandler).Methods("GET")
	r.HandleFunc("/mountains", mountainHandler).Methods("GET")
	r.HandleFunc("/health" , HealthCheck).Methods("GET")

	fmt.Println("Starting the server at 2050")

	http.ListenAndServe(":2050", r)
}

func healthHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Print("home handler")

}

func riverHandler (w http.ResponseWriter, r *http.Request) {
	data := []byte("River Handler called")

	writeJsonResponse(w, http.StatusOK, data)
}

func mountainHandler (w http.ResponseWriter, r *http.Request) {
	data := []byte("Mountain Handler called")
	writeJsonResponse(w, http.StatusOK, data)}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Since we're here, we already know that HTTP service is up. Let's just check the state of the boltdb connection

		data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
		writeJsonResponse(w, http.StatusOK, data)

}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}
type healthCheckResponse struct {
	Status string `json:"status"`
}