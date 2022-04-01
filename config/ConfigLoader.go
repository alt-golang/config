package config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type ConfigLoader struct {
	ConfigDir string
	Config    map[string]interface{}
}

func (configLoader ConfigLoader) LoadConfig() map[string]interface{} {
	return configLoader.LoadConfigWithDir(configLoader.ConfigDir)
}

func (configLoader ConfigLoader) LoadConfigWithDir(configDir string) map[string]interface{} {

	path := ""
	config := map[string]interface{}{
		//empty
	}

	if &configDir == nil || len(configDir) == 0 {
		path = configLoader.ConfigDir
	} else {
		path = configDir
	}

	if _, err := os.Stat(path); err == nil {
		config = configLoader.LoadConfigByPrecedence(path)
	}
	return config
}

func (configLoader ConfigLoader) LoadConfigByPrecedence(configDir string) map[string]interface{} {
	config := map[string]interface{}{}
	dirpath := configDir + string(os.PathSeparator)
	env := os.Getenv("GO_ENV")
	if len(env) == 0 {
		env = "local"
	}
	instance := os.Getenv("GO_APP_INSTANCE")
	if len(instance) == 0 {
		instance = "development"
	}
	profiles := strings.Split(os.Getenv("GO_PROFILES_ACTIVE"), ",")
	precedence := make([]string, 0)
	precedence = append(precedence,
		dirpath+"default.json",
		dirpath+"default.yml",
		dirpath+"default.yaml",
		dirpath+"application.json",
		dirpath+"application.yml",
		dirpath+"application.yaml",
		dirpath+env+".json",
		dirpath+env+".yml",
		dirpath+env+".yaml",
		dirpath+env+"-"+instance+".json",
		dirpath+env+"-"+instance+".yml",
		dirpath+env+"-"+instance+".yaml")

	profileFilenames := make([]string, len(profiles)*3)
	for i := 0; i < len(profiles); i++ {
		offset := i * 3
		profileFilenames[offset+0] = dirpath + "application-" + profiles[i] + ".json"
		profileFilenames[offset+1] = dirpath + "application-" + profiles[i] + ".yml"
		profileFilenames[offset+2] = dirpath + "application-" + profiles[i] + ".yaml"
	}
	precedence = append(precedence, profileFilenames...)

	for i := 0; i < len(precedence); i++ {
		filepath := precedence[i]
		precendentConfig := map[string]interface{}{}
		if _, err := os.Stat(filepath); err == nil {
			if strings.HasSuffix(filepath, ".json") {
				file, _ := os.Open(filepath)
				defer file.Close()
				decoder := json.NewDecoder(file)
				decoder.Decode(&precendentConfig)
			}
			if strings.HasSuffix(filepath, ".yml") || strings.HasSuffix(filepath, ".yaml") {
				if yamlString, err := os.ReadFile(filepath); err == nil {
					yaml.Unmarshal([]byte(string(yamlString)), &precendentConfig)
				}

			}
			for k, v := range precendentConfig {
				config[k] = v
			}
		}
	}

	return config
}
