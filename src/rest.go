package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter ,r *http.Request) {

	fmt.Println("Saying Hi")
	fmt.Fprintf(w, "Welcome Home!") // send data to client side
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc(/,homeHandler)
	fmt.Println("So far so good")
}
