//declaring the name of our package
package main

/* Importing packages
'net/http' - we need this package if we are to do some RESTful API-ing!
'log' will help us log an errors we need to */
import (
	"log"
	"net/http"
)

/* we declare a struct here called struct */
type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "You just hit my API!"}`))
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
