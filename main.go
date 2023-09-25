//This is the name of the package

package main

import (
	"fmt"      //Library for I/O
	"net/http" //for the creation of a web server

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")

	staticfileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/asetts/").Handler(staticfileHandler).Methods("GET")
	return r
}

func main() {

	r := NewRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
