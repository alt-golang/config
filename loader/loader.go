package loader

import (
	"alt-golang/config/properties"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

func LoadConfig() map[string]interface{} {
	return LoadConfigWithDir("config")
}

func LoadConfigWithDir(configDir string) map[string]interface{} {

	path := ""
	config := map[string]interface{}{
		//empty
	}

	if &configDir == nil || len(configDir) == 0 {
		path = "config"
	} else {
		path = configDir
	}

	if _, err := os.Stat(path); err == nil {
		config = LoadConfigByPrecedence(path)
	}
	return config
}

func LoadConfigByPrecedence(configDir string) map[string]interface{} {
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
		dirpath+"default.props",
		dirpath+"default.properties",
		dirpath+"application.json",
		dirpath+"application.yml",
		dirpath+"application.yaml",
		dirpath+"application.props",
		dirpath+"application.properties",
		dirpath+env+".json",
		dirpath+env+".yml",
		dirpath+env+".yaml",
		dirpath+env+".props",
		dirpath+env+".properties",
		dirpath+env+"-"+instance+".json",
		dirpath+env+"-"+instance+".yml",
		dirpath+env+"-"+instance+".yaml",
		dirpath+env+"-"+instance+".props",
		dirpath+env+"-"+instance+".properties")

	profileFilenames := make([]string, len(profiles)*3)
	for i := 0; i < len(profiles); i++ {
		offset := i * 3
		profileFilenames[offset+0] = dirpath + "application-" + profiles[i] + ".json"
		profileFilenames[offset+1] = dirpath + "application-" + profiles[i] + ".yml"
		profileFilenames[offset+2] = dirpath + "application-" + profiles[i] + ".yaml"
		profileFilenames[offset+2] = dirpath + "application-" + profiles[i] + ".props"
		profileFilenames[offset+2] = dirpath + "application-" + profiles[i] + ".properties"
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
			if strings.HasSuffix(filepath, ".props") || strings.HasSuffix(filepath, ".properties") {
				if propertiesString, err := os.ReadFile(filepath); err == nil {
					precendentConfig = properties.Parse(string(propertiesString))
				}
			}
			for k, v := range precendentConfig {
				config[k] = v
			}
		}
	}

	return config
}
