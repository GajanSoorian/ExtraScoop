package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ExtraScoop/src/rest"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rest.HomeHandler)

	servicesRouter := r.PathPrefix("/services").Subrouter()

	servicesRouter.HandleFunc("", rest.ServiceHandler)
	servicesRouter.Methods("GET").Path("/{serchcriteria}/{search}").HandlerFunc(rest.ServiceHandler)

	fmt.Println("So far so good")
	err := http.ListenAndServe(":9000", r) // set listen port. default handler DefaultServeMux called since nil is passed.
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
