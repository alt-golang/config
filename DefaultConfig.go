package config

import (
	"errors"
	"reflect"
	"strings"
)

type DefaultConfig struct {
	object map[string]interface{}
	path   string
	Config
}

func (config DefaultConfig) Get(path string) (interface{}, error) {
	return config.GetWithDefault(path, nil)
}

func (config DefaultConfig) GetWithDefault(path string, defaultValue interface{}) (interface{}, error) {
	if config.object[path] != nil {
		return config.object[path], nil
	}
	pathSteps := strings.Split(path, ".")
	root := config.object
	for i := 0; i < len(pathSteps) && root != nil; i++ {
		if root[pathSteps[i]] != nil {
			if reflect.TypeOf(root[pathSteps[i]]).Kind() == reflect.Map {
				root = root[pathSteps[i]].(map[string]interface{})
			} else {
				return root[pathSteps[i]], nil
			}
		} else {
			root = nil
		}
	}
	if root != nil {
		return root, nil
	}
	if defaultValue != nil {
		return defaultValue, nil
	}
	return nil, errors.New("DefaultConfig path " + config.path + " returned no value.")
}

func (config DefaultConfig) Has(path string) bool {
	if config.object[path] != nil {
		return true
	}
	pathSteps := strings.Split(path, ".")
	root := config.object
	for i := 0; i < len(pathSteps) && root != nil; i++ {
		if root[pathSteps[i]] != nil {
			if reflect.TypeOf(root[pathSteps[i]]).Kind() == reflect.Map {
				root = root[pathSteps[i]].(map[string]interface{})
			} else {
				return true
			}
		} else {
			root = nil
		}

	}
	return root != nil
}
