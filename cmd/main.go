package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ZAF07/cost-tracker/api/rest/router"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server")
	r := mux.NewRouter()
	router.InitRoutes(r)

	// log.Fatal(http.ListenAndServe(":8080", r))
	if err := http.ListenAndServe(":8000", r); err != nil {
		// log error here
		log.Fatal(err)
	}
}

// SET UP PROPER LOGGING SYSTEM
// SET UP FIRST SERVICE
// SET UP FIRST CONTROLLER
// SET UP BASIC ROUTER
//  SET UP BASIC GORILLA SERVER