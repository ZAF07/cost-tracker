package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type AboutAPI struct {
	// db *gorm.DB
}

func NewAboutAPI() *AboutAPI {
	return&AboutAPI{}
}

type AboutResponse struct {
	Page string
}

func (c AboutAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello %s", "ZAFFERE")

		h := w.Header()
		h.Add("token", "1234")
	
		t := AboutResponse{
			Page: "About Page",
		}
		response, marshalErr := json.Marshal(t)
	
		if marshalErr != nil {
			// LOG ERROR HERE
			log.Fatal(marshalErr)
		}
		w.Write(response)
}