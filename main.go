package main

import (
	"encoding/json"
	"fmt"
	"go-platform-service/com/nike/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", healthHandler).Methods("GET")
	r.HandleFunc("/rivers", riverHandler).Methods("GET")
	r.HandleFunc("/mountains", mountainHandler).Methods("GET")
	r.HandleFunc("/platform/states", statesHandler).Methods("GET")

	r.HandleFunc("/health", HealthCheck).Methods("GET")

	fmt.Println("Starting the server at 8080")

	_ = http.ListenAndServe(":8080", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("home handler")
}

func riverHandler(w http.ResponseWriter, r *http.Request) {
	data := []byte("River Handler called")
	writeJsonResponse(w, http.StatusOK, data)
}

func mountainHandler(w http.ResponseWriter, r *http.Request) {
	data := []byte("Mountain Handler called")
	writeJsonResponse(w, http.StatusOK, data)
}

func statesHandler(w http.ResponseWriter, r *http.Request) {

	response := service.InvokeStateService("Oregon")
	data := []byte(response)
	writeJsonResponse(w, http.StatusOK, data)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Since we're here, we already know that HTTP service is up. Let's just check the state of the boltdb connection
	data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
	writeJsonResponse(w, http.StatusOK, data)
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

type healthCheckResponse struct {
	Status string `json:"status"`
}
