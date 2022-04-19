package loader

import (
	"encoding/json"
	"github.com/alt-golang/config/properties"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
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
	precedence := make([]interface{}, 0)
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

	profileFilenames := make([]interface{}, len(profiles)*5)
	for i := 0; i < len(profiles); i++ {
		offset := i * 5
		profileFilenames[offset+0] = dirpath + "application-" + profiles[i] + ".json"
		profileFilenames[offset+1] = dirpath + "application-" + profiles[i] + ".yml"
		profileFilenames[offset+2] = dirpath + "application-" + profiles[i] + ".yaml"
		profileFilenames[offset+3] = dirpath + "application-" + profiles[i] + ".props"
		profileFilenames[offset+4] = dirpath + "application-" + profiles[i] + ".properties"
	}
	precedence = append(precedence, profileFilenames...)

	for i := 0; i < len(precedence); i++ {
		filepath := precedence[i].(string)
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
					yamlConfig := map[interface{}]interface{}{}
					yaml.Unmarshal([]byte(string(yamlString)), &yamlConfig)
					precendentConfig = Normalise(yamlConfig)
				}
			}
			if strings.HasSuffix(filepath, ".props") || strings.HasSuffix(filepath, ".properties") {
				if propertiesString, err := os.ReadFile(filepath); err == nil {
					precendentConfig = map[string]interface{}(properties.Parse(string(propertiesString)))
				}
			}
			AssignIn(config, precendentConfig)
		}
	}
	envvars := map[string]interface{}{}
	for _, env := range os.Environ() {
		envPair := strings.SplitN(env, "=", 2)
		key := envPair[0]
		value := envPair[1]
		envvars[key] = value
	}
	config["env"] = envvars
	args := make([]string, 0)
	config["args"] = append(args, os.Args...)

	return config
}

func Normalise(object map[interface{}]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for key, value := range object {
		if reflect.TypeOf(value) == reflect.TypeOf(map[interface{}]interface{}{}) {
			result[key.(string)] = Normalise(value.(map[interface{}]interface{}))
		} else if reflect.TypeOf(value).Kind() == reflect.Slice {
			s := value.([]interface{})
			for i := 0; i < len(s); i++ {
				if reflect.TypeOf(s[i]) == reflect.TypeOf(map[interface{}]interface{}{}) {
					s[i] = Normalise(s[i].(map[interface{}]interface{}))
				}
			}
			result[key.(string)] = s
		} else {
			result[key.(string)] = value
		}
	}
	return result
}

func AssignIn(config map[string]interface{}, object map[string]interface{}) {

	for i := range object {
		key := i

		if config[key] == nil {
			config[key] = object[key]
		} else {
			if reflect.TypeOf(config[key]) == reflect.TypeOf(map[string]interface{}{}) {
				if reflect.TypeOf(object[key]) == reflect.TypeOf(map[string]interface{}{}) {
					AssignIn(config[key].(map[string]interface{}), object[key].(map[string]interface{}))
				} else {
					config[key] = object[key]
				}
			} else if reflect.TypeOf(config[key]).Kind() == reflect.Slice &&
				reflect.TypeOf(object[key]).Kind() == reflect.Slice {
				configList := config[key].([]interface{})
				objectList := object[key].([]interface{})
				for j := 0; j < len(objectList); j++ {
					if j < len(configList) {
						if reflect.TypeOf(configList[j]) == reflect.TypeOf(map[string]interface{}{}) {
							if reflect.TypeOf(objectList[j]) == reflect.TypeOf(map[string]interface{}{}) {
								AssignIn(configList[j].(map[string]interface{}), objectList[j].(map[string]interface{}))
							} else {
								configList[j] = objectList[j]
							}
						} else {
							configList[j] = objectList[j]
						}
					} else {
						config[key] = append(configList, objectList[j])
					}
				}
			} else {
				config[key] = object[key]
			}
		}
	}
}
