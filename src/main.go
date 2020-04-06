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
	//handler = new()
	r.HandleFunc("/", rest.HomeHandler)

	servicesRouter := r.PathPrefix("/services").Subrouter()
	//test start
	servicesRouter.HandleFunc("", rest.TestServiceHandler)
	servicesRouter.Methods("GET").Path("/{serchcriteria}/{search}").HandlerFunc(rest.TestServiceHandler)
	//test end
	servicesRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(rest.FindServiceHandler)
	servicesRouter.Methods("GET").Path("").HandlerFunc(rest.AllServiceHandler)
	servicesRouter.Methods("POST").Path("").HandlerFunc(rest.NewServiceHandler)

	fmt.Println("So far so good")
	err := http.ListenAndServe(":9000", r) // set listen port. default handler DefaultServeMux called since nil is passed.
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
