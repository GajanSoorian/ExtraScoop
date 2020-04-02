package rest

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Saying Hi")
	fmt.Fprintf(w, "Welcome Home!") // send data to client side
}

func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("using Service")
	fmt.Fprintf(w, "These are the list of services!\n")       // send data to client side
	fmt.Fprintf(w, "1) Hair cut \n 2) Shave \n 3) manicure ") // send data to client side
}
