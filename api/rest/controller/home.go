package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HomeAPI struct {
	// db *gorm.DB
}

func NewAppAPI() *HomeAPI {
	return &HomeAPI{}
}

type Response struct {
	Page string
}

func (c HomeAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	func(w http.ResponseWriter, r *http.Request) {
		// Bases on method, render the relevant service module, passing the app dependencies (DB, RMQ, CACHE)
		fmt.Println("REQUEST METHOD: ", r.Method)
		h := w.Header()
		h.Add("token", "1234")

		t := Response{
			Page: "Home Page",
		}
		response, marshalErr := json.Marshal(t)
		if marshalErr != nil {
			// LOG ERROR HERE
			log.Fatal(marshalErr)
		}

		w.Write(response)
	}(w, r)
}
