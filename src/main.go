package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ExtraScoop/src/persistence/dblayer"
	"github.com/ExtraScoop/src/rest"
	"github.com/gorilla/mux"
)

func startServer(endpoint string) {
	r := mux.NewRouter()
	handler := &rest.ServicesHandler{}
	r.HandleFunc("/", rest.HomeHandler)

	servicesRouter := r.PathPrefix("/services").Subrouter()
	//test start
	servicesRouter.HandleFunc("", rest.TestServiceHandler)
	servicesRouter.Methods("GET").Path("/{serchcriteria}/{search}").HandlerFunc(rest.TestServiceHandler)
	//test end
	servicesRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindServiceHandler)
	servicesRouter.Methods("GET").Path("").HandlerFunc(handler.AllServiceHandler)
	servicesRouter.Methods("POST").Path("").HandlerFunc(handler.NewServiceHandler)

	fmt.Println("So far so good")
	err := http.ListenAndServe(endpoint, r) // set listen port. default handler DefaultServeMux called since nil is passed.
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer("mongodb", "mongodb://localhost:27017")
	dbhandler.AddService()
	startServer(":9000")
}
