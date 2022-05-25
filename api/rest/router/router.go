package router

import (
	"fmt"
	"log"

	"github.com/ZAF07/cost-tracker/internal/config"
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router, paths config.AppPaths) {
	if len(paths.Paths) < 1 {
		// USE GLOBAL LOGGER
		log.Fatal("LESS THAN ONE")
	}

	for _, p := range paths.Paths {
		methods := []string{}
		for _, v := range p.Methods {
			methods = append(methods, v.URL)
		}
		fmt.Println()
		fmt.Printf("REGISTERING THESE METHODS FOR /%s : %v", p.Path, methods)
		fmt.Println("")
		r.HandleFunc(p.URL, p.Handle.ServeHTTP).Methods(methods...)
		// Figure out timeout
	}
}
