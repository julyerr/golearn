package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	v := router.PathPrefix("/hello").Subrouter()
	v.HandleFunc("/say", sayHello)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
