//This is the name of the package

package main

import (
	"fmt"      //Library for I/O
	"net/http" //for the creation of a web server

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	//This is a constructor function It creates a router object using gorilla

	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	//This is the file directory in html
	staticFileDirectory := http.Dir("./assets/")

	staticfileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/assets/").Handler(staticfileHandler).Methods("GET")
	return r
}

func main() {

	//This just creates a router function and
	//listens and serve in the 8080 port.
	r := NewRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
