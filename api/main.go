package main

import (
	"encoding/json"
	"log"
	"net/http"
)
// allow CORS

func main() {
	http.HandleFunc("/data", getData)
	log.Fatal(http.ListenAndServe(":8080", addCorsHeaders(http.DefaultServeMux)))
}

func getData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Prepare the response data
	data := []int{10, 20, 30, 40, 50}
	labels := []string{"January", "February", "March", "April", "May"}

	// Prepare the response
	response := struct {
		Data []int `json:"data"`
		Labels []string `json:"labels"`
	}{
		Data: data,
		Labels: labels,
	}

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
func addCorsHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow the GET method
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		// Allow the Content-Type header
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Continue handling the request
		handler.ServeHTTP(w, r)
	})
}