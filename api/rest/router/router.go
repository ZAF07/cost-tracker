package router

import (
	"errors"
	"fmt"
	"log"

	"github.com/ZAF07/cost-tracker/internal/config"
	"github.com/gorilla/mux"
)

func NewRouter(p config.AppPaths) (r *mux.Router, err error) {
	r = mux.NewRouter()
	if err := initPaths(r, p); err != nil {
		return r, err
	}
	return
}

func initPaths(r *mux.Router, paths config.AppPaths) error {
	if len(paths.Paths) < 1 {
		// USE GLOBAL LOGGER
		log.Fatal("LESS THAN ONE")
		return errors.New("invalid paths")
	}

	for _, p := range paths.Paths {
		methods := []string{}
		for _, v := range p.Methods {
			methods = append(methods, v.URL)
		}
		fmt.Printf("REGISTERING THESE METHODS FOR /%s : %v\n", p.Path, methods)
		r.HandleFunc(p.URL, p.Handle.ServeHTTP).Methods(methods...)
		// Figure out timeout
	}
	return nil
}
