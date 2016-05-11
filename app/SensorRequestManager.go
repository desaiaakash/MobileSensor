package main

import (
	
	"github.com/julienschmidt/httprouter"
	"net/http"
	)

	func routeRequest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	
	}


	func main() {

	mux := httprouter.New()
	
	mux.POST("/", routeRequest)			      					// Routing the request to any one of two stations
	
	
	
	server := http.Server{
		Addr:    "0.0.0.0:3500",
		Handler: mux,
	}
	server.ListenAndServe()

}

