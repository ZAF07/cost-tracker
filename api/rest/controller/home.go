package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type HomeAPI struct {
	// db *gorm.DB
}

func NewAppAPI() *HomeAPI {
	return&HomeAPI{}
}

type Response struct {
	Page string
}

func (c HomeAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello %s", "ZAFFERE")
	func (w http.ResponseWriter, r *http.Request) {

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
	}(w,r)
}

