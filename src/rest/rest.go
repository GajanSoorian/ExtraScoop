package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Saying Hi")
	fmt.Fprintf(w, "Welcome Home!") // send data to client side
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("using Service")
	fmt.Fprintf(w, "These are the list of services!\n")       // send data to client side
	fmt.Fprintf(w, "1) Hair cut \n 2) Shave \n 3) manicure ") // send data to client side
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	servicesRouter := r.PathPrefix("/services").Subrouter()

	servicesRouter.HandleFunc("/", serviceHandler)
	fmt.Println("So far so good")
	err := http.ListenAndServe(":9000", r) // set listen port. default handler DefaultServeMux called since nil is passed.
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
