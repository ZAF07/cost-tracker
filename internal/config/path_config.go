package config

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ZAF07/cost-tracker/api/rest/controller"
)

type AppPaths struct {
	Paths []Path `json:"paths"`
}

type Path struct {
	Path string `json:"path"`
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
						p.Handle = controller.MapControllerHandler(value.(string))
						return nil
					}
					p.Handle = controller.NewAboutAPI()
			}
	}
	return nil
}