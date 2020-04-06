package rest

import (
	"fmt"
	"net/http"
)

type servicesHandler struct {
}

//HomeHandler is a handler for home page test
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Saying Hi")
	fmt.Fprintf(w, "Welcome Home!") // send data to client side
}

//TestServiceHandler is a handler for testing service
func TestServiceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("using Service")
	fmt.Fprintf(w, "These are the list of services!\n")       // send data to client side
	fmt.Fprintf(w, "1) Hair cut \n 2) Shave \n 3) manicure ") // send data to client side
}

func /*(sh *servicesHandler)*/ AllServiceHandler(w http.ResponseWriter, r *http.Request) {
}

func /*(sh *servicesHandler)*/ NewServiceHandler(w http.ResponseWriter, r *http.Request) {
}

func /*(sh *servicesHandler)*/ FindServiceHandler(w http.ResponseWriter, r *http.Request) {
}
