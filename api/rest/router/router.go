package router

import (
	"fmt"
	"log"

	"github.com/ZAF07/cost-tracker/internal/config"
	"github.com/gorilla/mux"
)

// GET THIS FROM CONFIG FILE
var appPaths = `
{
  "paths": [
    {
      "url": "/about",
      "handler": "about"
    },
    {
      "url": "/home",
      "handler": "home"
    }
  ]
}
`
var GET string = ""
var POST string = ""
var PUT string = ""
var DELETE string = ""

func InitRoutes(r *mux.Router, paths config.AppPaths) {
	if len(paths.Paths) < 1 {
		log.Fatal("LESS THAN ONE")
	}

	for _, p := range paths.Paths {

		// a := ""
		for _, v := range p.Methods {
			switch v.URL {
			case "GET":
				GET = v.URL
			case "POST":
				POST = v.URL
			case "PUT":
				POST = v.URL
			case "DELETE":
				DELETE = v.URL
			}
		}
		fmt.Println()
		fmt.Printf("REGISTERING THESE METHODS FOR /%s : %s %s %s %s", p.Path, GET, POST, PUT, DELETE)
		fmt.Println("")
		r.HandleFunc(p.URL, p.Handle.ServeHTTP).Methods(GET, POST, PUT, DELETE)
		// fmt.Println(r)
	}
}
