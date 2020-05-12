package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	//read something
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET request received"`))
	//create something
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST request received"`))
	//modify something
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "PUT request received"`))
	//remove something
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "DELETE request received"`))
	//catch-all
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "No Found"`))
	}
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
