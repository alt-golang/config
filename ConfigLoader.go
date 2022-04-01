package main

import (
	"os"
)

type ConfigLoader struct {
	ConfigDir string
	Config    map[string]interface{}
}

func (configLoader ConfigLoader) LoadConfig() map[string]interface{} {
	configLoader.LoadConfigWithDir(configLoader.ConfigDir)
}

func (configLoader ConfigLoader) LoadConfigWithDir(configDir string) map[string]interface{} {

	var path string

	if &configDir == nil || len(configDir) == 0 {
		path = configLoader.ConfigDir
	} else {
		path = configDir
	}

	if _, err := os.Stat(path); err == nil {
		return configLoader.LoadConfigByPrecedence(path)
	}

	var config map[string]interface{}
	return config
}

func (configLoader ConfigLoader) LoadConfigByPrecedence(configDir string) map[string]interface{} {
	return nil
}
