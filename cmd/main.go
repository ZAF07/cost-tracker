package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ZAF07/cost-tracker/api/rest/controller"
	"github.com/ZAF07/cost-tracker/api/rest/router"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Paths struct {
	Paths []Path `json:"paths"`
}
type Path struct {
	Path string `json:"path"`
	URL string `json:"url"`
	Handler string `json:"handler"`
	Handle controller.Con
}

//  Custom Unmarshal to init controller struct
func (p *Path) UnmarshalYAML(data []byte) error {
	 var v map[string]interface{}

		err := json.Unmarshal(data, &v)
		if err != nil {
			fmt.Printf("%+v", data)
			log.Fatal("PATH CUSTOM UNMARSHAL: ", err)
			return err
		}

		log.Printf("HERE ----- : %+v", v)
		for key, value := range v {
			switch key {
				case "url":
					p.URL = value.(string)
				case "handler":
					p.Handler = value.(string)
					if value == "home" {
						p.Handle = controller.NewAppAPI()
						return nil
					}
					p.Handle = controller.NewAboutAPI()
			}
	}
	return nil
}
func main() {
	fmt.Println("Server")
	r := mux.NewRouter()
	router.InitRoutes(r)

    viper.AddConfigPath("config")
    viper.SetConfigName("config")
    viper.SetConfigType("yml")
		
		s := Paths{}
    err := viper.ReadInConfig()
		erre := viper.Unmarshal(&s)
		if erre != nil {
			log.Fatalf("UNMARSHAL_JSON: ", erre)
		}
		if err != nil {
			log.Fatalf("Error reading config file; ", err)
		}
		
    f := viper.Get("paths")
		fmt.Println()
    fmt.Println("Result config: ", f)
		// Option 1: Loop through the unmarshalled data and and init a new controller based on its handler
		s.Paths[0].Handle = controller.NewAppAPI()
		fmt.Printf("THIS IS THE STRUCT : %+v", s.Paths[0])



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