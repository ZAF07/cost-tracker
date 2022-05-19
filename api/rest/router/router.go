package router

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ZAF07/cost-tracker/api/rest/controller"
	"github.com/gorilla/mux"
)

type Paths struct {
	Paths []Path `json:"paths"`
}
type Path struct {
	URL string `json:"url"`
	Handler string `json:"handler"`
	Handle controller.Con
}

//  Custom Unmarshal to init controller struct
func (p *Path) UnmarshalJSON(data []byte) error {
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
func InitRoutes(r *mux.Router){
	ps := &Paths{}
	err := json.Unmarshal([]byte(appPaths), ps)
	if err != nil {
		log.Fatal("INITROUTES: ",err)
	}

	if len( ps.Paths) < 1 {
		log.Fatal("LESS THAN ONE")
	}
	for _, p := range ps.Paths {
		fmt.Printf("%+v", p)
		r.HandleFunc(p.URL, p.Handle.ServeHTTP)
	}
}