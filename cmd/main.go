package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if json, err := json.Marshal("Hello world."); err == nil {
			fmt.Fprintf(w, string(json))
		}  else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}).Methods("GET")

	return r
}

func main() {
	fmt.Printf("Server started on port 8080...\n")
	http.ListenAndServe(":8080", newRouter())
}
