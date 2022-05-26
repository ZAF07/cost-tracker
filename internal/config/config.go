package config

import (
	"fmt"
	"log"
	"reflect"

	"github.com/ZAF07/cost-tracker/api/rest/controller"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type GeneralConfig AppConfig

func InitialisePaths() (p AppPaths, err error) {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	s := AppPaths{}
	readConfigErr := viper.ReadInConfig()
	if readConfigErr != nil {
		log.Fatalf("Error reading config file; %+v", err)
		return AppPaths{}, readConfigErr
	}
	// INIT AppConfig{}

	erre := viper.Unmarshal(&s, viper.DecodeHook(UnmarshalHookYAML()))
	if erre != nil {
		log.Fatalf("UNMARSHAL_JSON: %+v", erre)
		return AppPaths{}, readConfigErr
	}
	return s, nil
}

func UnmarshalHookYAML() mapstructure.DecodeHookFuncType {
	// Wrapped in a function call to add optional input parameters (eg. separator)
	return func(
		f reflect.Type, // data type
		t reflect.Type, // target data type
		data interface{}, // raw data
	) (interface{}, error) {
		// Check if the target type matches Path{}
		appPath := Path{}
		if t != reflect.TypeOf(appPath) {
			return data, nil
		}

		fm := reflect.ValueOf(data)

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
				methds := []Method{}
				for i := 0; i < int(values.Int()); i++ {
					mtd := Method{}
					// mtd.URL = iSlice.Index(i).String().(string)
					//  Changed to convert to interface so that we can assert to string
					mtd.URL = iSlice.Index(i).Interface().(string)
					methds = append(methds, mtd)
				}
				appPath.Methods = methds

			}

			switch lo.Key().Interface().(string) {
			case "url":
				appPath.URL = v
			case "handler":
				appPath.Handler = v
				appPath.Handle = controller.MapControllerHandler(v)
			case "path":
				appPath.Path = v
			}
		}
		return appPath, nil
	}
}
