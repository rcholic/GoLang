package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

//ServeHTTP on String type
func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

//ServeHTTP handler on Struct
func (m Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, m.Greeting, m.Punct, m.Who)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
