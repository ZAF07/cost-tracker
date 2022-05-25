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
	Method  string `mapstructure:"method" json:"method"`
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
		pa := config.Path{}

		lo := fm.MapRange()
		for lo.Next() {
			// methodSlice := []interface{}{}
			v, ok := lo.Value().Interface().(string)
			if !ok {
				//  Have to call the Elem() to return the interface's true value
				iSlice := reflect.Value(lo.Value().Elem())
				fmt.Println("ISLICE : ", iSlice)

				values := reflect.ValueOf(lo.Value().Elem().Len())
				fmt.Println(values)
				methds := []config.Method{}
				for i := 0; i < int(values.Int()); i++ {
					mtd := config.Method{}
					// mtd.URL = iSlice.Index(i).String().(string)
					//  Changed to convert to interface so that we can assert to string
					mtd.URL = iSlice.Index(i).Interface().(string)
					methds = append(methds, mtd)
				}
				pa.Methods = methds

			}

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

	// THIS GOES INTO ROUTER PACKAGE
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
//  SET UP BASIC MUX SERVER
