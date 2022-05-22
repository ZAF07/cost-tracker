package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/ZAF07/cost-tracker/api/rest/controller"
	"github.com/ZAF07/cost-tracker/api/rest/router"
	"github.com/ZAF07/cost-tracker/internal/config"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Paths struct {
	Paths []Path `json:"paths"`
}
type Path struct {
	Path    string `mapstructure:"path" json:"path"`
	URL     string `mapstructure:"url" json:"url"`
	Handler string `mapstructure:"handler" json:"handler"`
	Handle  controller.Con
}

func UnmarshalHookYAML() mapstructure.DecodeHookFuncType {
	// Wrapped in a function call to add optional input parameters (eg. separator)
	return func(
		f reflect.Type, // data type
		t reflect.Type, // target data type
		data interface{}, // raw data
	) (interface{}, error) {
		// Check if the target type matches Path{}
		configPaths := config.Path{}
		if t != reflect.TypeOf(configPaths) {
			return data, nil
		}

		fm := reflect.ValueOf(data)
		pa := Path{}

		lo := fm.MapRange()
		for lo.Next() {
			v := lo.Value().Interface().(string)
			fm := reflect.ValueOf(lo.Key()).Interface()
			log.Printf("THIS IS K : %+v", fm)
			switch lo.Key().Interface().(string) {
			case "url":
				pa.URL = v
			case "handler":
				pa.Handler = v
				pa.Handle = controller.MapControllerHandler(v)
			case "path":
				pa.Path = v
			}
		}
		return pa, nil
	}
}

func main() {
	fmt.Println("Server")
	r := mux.NewRouter()

	// THIS WOULD BE THE UNMARSHALER FOR ANY YAML CONFIG FILE
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	s := config.AppPaths{}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file; %+v", err)
	}
	erre := viper.Unmarshal(&s, viper.DecodeHook(UnmarshalHookYAML()))
	if erre != nil {
		log.Fatalf("UNMARSHAL_JSON: %+v", erre)
	}
	fmt.Printf("THIS IS THE FINAL STRUCT : %+v", s.Paths)
	router.InitRoutes(r, s)

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
