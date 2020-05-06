package main

import (
	"fmt"
	"log"
	"net/http"
)

func testAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You found me!")
	fmt.Println("This is the EndPoint for my Go API")
}

func handleRequests() {
	http.HandleFunc("/", testAPI)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
