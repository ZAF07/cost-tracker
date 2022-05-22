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

func InitRoutes(r *mux.Router, paths config.AppPaths) {
	if len(paths.Paths) < 1 {
		log.Fatal("LESS THAN ONE")
	}
	for _, p := range paths.Paths {
		fmt.Printf("%+v", p)
		r.HandleFunc(p.URL, p.Handle.ServeHTTP)
	}
}
