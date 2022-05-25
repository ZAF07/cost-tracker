package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ZAF07/cost-tracker/api/rest/router"
	"github.com/ZAF07/cost-tracker/internal/config"
)

func main() {
	fmt.Println("Server")

	paths, err := config.InitialisePaths()
	if err != nil {
		log.Panicf("error initialising paths: %v", err)
	}
	appRouter, err := router.NewRouter(paths)
	if err != nil {
		log.Panicf("Error initialising router in main %+v", err)
	}

	if err := http.ListenAndServe(":8000", appRouter); err != nil {
		// global log error here
		log.Fatal(err)
	}
}

// SET UP PROPER LOGGING SYSTEM
// SET UP FIRST SERVICE
// SET UP FIRST CONTROLLER
// SET UP BASIC ROUTER
//  SET UP BASIC MUX SERVER
