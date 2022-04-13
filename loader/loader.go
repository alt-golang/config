package loader

import (
	"alt-golang/config/properties"
	"encoding/json"
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

	return config
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
				} else if reflect.TypeOf(object[key]) == reflect.TypeOf(map[interface{}]interface{}{}) {
					AssignIn(config[key].(map[string]interface{}), ConvertIn(object[key].(map[interface{}]interface{})))
				} else {
					config[key] = object[key]
				}

			} else if reflect.TypeOf(config[key]) == reflect.TypeOf(map[interface{}]interface{}{}) {
				if reflect.TypeOf(object[key]) == reflect.TypeOf(map[string]interface{}{}) {
					AssignIn(ConvertIn(config[key].(map[interface{}]interface{})), object[key].(map[string]interface{}))
				} else if reflect.TypeOf(object[key]) == reflect.TypeOf(map[interface{}]interface{}{}) {
					AssignIn(ConvertIn(config[key].(map[interface{}]interface{})), ConvertIn(object[key].(map[interface{}]interface{})))
				} else {
					config[key] = object[key]
				}
			} else if reflect.TypeOf(config[key]) == reflect.TypeOf([]interface{}{}) &&
				reflect.TypeOf(object[key]) == reflect.TypeOf([]interface{}{}) {
				configList := config[key].([]interface{})
				objectList := object[key].([]interface{})
				for j := 0; j < len(objectList); j++ {
					if j < len(configList) {
						if reflect.TypeOf(configList[j]) == reflect.TypeOf(map[string]interface{}{}) {
							if reflect.TypeOf(objectList[j]) == reflect.TypeOf(map[string]interface{}{}) {
								AssignIn(configList[j].(map[string]interface{}), objectList[j].(map[string]interface{}))
							} else if reflect.TypeOf(objectList[j]) == reflect.TypeOf(map[interface{}]interface{}{}) {
								AssignIn(configList[j].(map[string]interface{}), ConvertIn(objectList[j].(map[interface{}]interface{})))
							} else {
								configList[j] = objectList[j]
							}
						} else if reflect.TypeOf(configList[j]) == reflect.TypeOf(map[interface{}]interface{}{}) {
							if reflect.TypeOf(objectList[j]) == reflect.TypeOf(map[string]interface{}{}) {
								AssignIn(ConvertIn(configList[j].(map[interface{}]interface{})), objectList[j].(map[string]interface{}))
							} else if reflect.TypeOf(objectList[j]) == reflect.TypeOf(map[interface{}]interface{}{}) {
								AssignIn(ConvertIn(configList[j].(map[interface{}]interface{})), ConvertIn(objectList[j].(map[interface{}]interface{})))
							} else {
								configList[j] = objectList[j]
							}
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

func ConvertIn(object map[interface{}]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for key, value := range object {
		result[key.(string)] = value
	}
	return result
}

func Normalise(object map[interface{}]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for key, value := range object {
		if reflect.TypeOf(value) == reflect.TypeOf(map[interface{}]interface{}{}) {
			result[key.(string)] = Normalise(value.(map[interface{}]interface{}))
		} else {
			result[key.(string)] = value
		}
	}
	return result
}

/*
 assignIn(config, object) {
    const keys = Object.keys(object);
    for (let i = 0; i < keys.length; i++) {
      const key = keys[i];
      if (typeof (config[key]) === 'undefined') {
        // eslint-disable-next-line   no-param-reassign
        config[key] = object[key];
      } else if (typeof (config[key]) === 'object') {
        if (config[key]
        && typeof (config[key]) !== 'undefined'
            && typeof (config[key]) === 'object'
            && typeof (object[key]) === 'object') {
          this.assignIn(config[key], object[key]);
        } else if (Array.isArray(config[key]) && Array.isArray(object[key])) {
          for (let j = 0; j < object[key].length; j++) {
            if (j < config[key].length) {
              this.assignIn(config[key][j], object[key][j]);
            } else {
              config[key].push(object[key][j]);
            }
          }
        } else {
          // eslint-disable-next-line   no-param-reassign
          config[key] = object[key];
        }
      } else {
        // eslint-disable-next-line   no-param-reassign
        config[key] = object[key];
      }
    }
  }
*/
