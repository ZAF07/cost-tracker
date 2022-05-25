package router

import (
	"fmt"
	"log"

	"github.com/ZAF07/cost-tracker/constants"
	"github.com/ZAF07/cost-tracker/internal/config"
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router, paths config.AppPaths) {
	if len(paths.Paths) < 1 {
		// USE GLOBAL LOGGER
		log.Fatal("LESS THAN ONE")
	}

	for _, p := range paths.Paths {
		var GET string = ""
		var POST string = ""
		var PUT string = ""
		var DELETE string = ""

		// a := ""
		for _, v := range p.Methods {
			switch v.URL {
			case constants.GET:
				GET = v.URL
			case constants.POST:
				POST = v.URL
			case constants.PUT:
				PUT = v.URL
			case constants.DELETE:
				DELETE = v.URL
			}
		}
		fmt.Println()
		fmt.Printf("REGISTERING THESE METHODS FOR /%s : %s %s %s %s", p.Path, GET, POST, PUT, DELETE)
		fmt.Println("")
		r.HandleFunc(p.URL, p.Handle.ServeHTTP).Methods(GET, POST, PUT, DELETE)
	}
}
