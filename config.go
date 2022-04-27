package config

import (
	"encoding/json"
	"github.com/alt-golang/config/loader"
)

type Config interface {
	Has(path string) bool
	Get(path string) (interface{}, error)
	GetAs(path string, object any) (any, error)
	GetWithDefault(path string, defaultValue interface{}) (interface{}, error)
}

var config = Factory(loader.LoadConfig())

func GetConfig() Config {
	return config
}

func GetConfigFromDir(dir string) Config {
	config = Factory(loader.LoadConfigWithDir(dir, false))
	return config
}

func GetServiceConfigFromDir(dir string) Config {
	config = ServiceFactory(loader.LoadConfigWithDir(dir, true))
	return config
}

func Has(path string) bool {
	return config.Has(path)
}

func Get(path string) (interface{}, error) {
	return config.Get(path)
}

func GetAs(path string, object any) error {
	config, err := config.Get(path)
	if err != nil {
		AssignAs(config.(map[string]any), object)
		return nil
	}
	return err
}

func GetWithDefault(path string, defaultValue interface{}) (interface{}, error) {
	return config.GetWithDefault(path, defaultValue)
}

func AssignAs(config map[string]interface{}, object interface{}) {
	if object != nil {
		result, _ := json.Marshal(config)
		json.Unmarshal([]byte(result), object)
	}
}
