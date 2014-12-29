package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string
type Struct struct {
	Greeting, Punct, Who string
}

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	texto := "OK string " + s + "!"
	fmt.Fprint(w, texto)
}

func (s *Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	texto := "OK struc " + s.Greeting + " " + s.Punct + " " + s.Who + "!"
	fmt.Fprint(w, texto)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
