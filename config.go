package config

import "github.com/alt-golang/config/loader"

type Config interface {
	Has(path string) bool
	Get(path string) (interface{}, error)
	GetWithDefault(path string, defaultValue interface{}) (interface{}, error)
}

var config = GetConfig()

func GetConfig() Config {
	return Factory(loader.LoadConfig())
}

func GetConfigWithDir(dir string) Config {
	return Factory(loader.LoadConfigWithDir(dir))
}

func Has(path string) bool {
	return config.Has(path)
}

func Get(path string) (interface{}, error) {
	return config.Get(path)
}

func GetWithDefault(path string, defaultValue interface{}) (interface{}, error) {
	return config.GetWithDefault(path, defaultValue)
}
