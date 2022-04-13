package loader

import (
	"os"
	"testing"
)

func TestJSONConfigIsOverloaded(t *testing.T) {

	os.Setenv("GO_ENV", "environment")
	os.Setenv("GO_APP_INSTANCE", "instance")
	os.Setenv("GO_PROFILES_ACTIVE", "1,2")
	var config map[string]interface{}
	config = LoadConfigWithDir("../test/config/json")

	if config["default"] != "default" {
		t.Errorf("config[\"default\"] != \"default\": default is:%s", config["default"])
	}
	if config["application"] != "application" {
		t.Errorf("config[\"application\"] != \"application\": application is:%s", config["application"])
	}
	if config["environment"] != "environment" {
		t.Errorf("config[\"environment\"] != \"environment\": environment is:%s", config["environment"])
	}
	if config["env-instance"] != "env-instance" {
		t.Errorf("config[\"env-instance\"] != \"env-instance\": env-instance is:%s", config["environment"])
	}
	if config["local"] != "default" {
		t.Errorf("config[\"local\"] != \"default\": local is:%s", config["local"])
	}

	os.Setenv("GO_ENV", "")
	os.Setenv("GO_APP_INSTANCE", "")
	os.Setenv("GO_PROFILES_ACTIVE", "")
	config = LoadConfigWithDir("../test/config/json")
	if config["local"] != "local" {
		t.Errorf("config[\"local\"] != \"local\": local is:%s", config["local"])
	}
	if config["local-development"] != "local-development" {
		t.Errorf("config[\"local-development\"] != \"local-development\": local-development is:%s", config["local-development"])
	}

}

func TestYAMLConfigIsOverloaded(t *testing.T) {

	os.Setenv("GO_ENV", "environment")
	os.Setenv("GO_APP_INSTANCE", "instance")
	os.Setenv("GO_PROFILES_ACTIVE", "1,2")
	var config map[string]interface{}
	config = LoadConfigWithDir("../test/config/yaml")

	if config["default"] != "default" {
		t.Errorf("config[\"default\"] != \"default\": default is:%s", config["default"])
	}
	if config["application"] != "application" {
		t.Errorf("config[\"application\"] != \"application\": application is:%s", config["application"])
	}
	if config["environment"] != "environment" {
		t.Errorf("config[\"environment\"] != \"environment\": environment is:%s", config["environment"])
	}
	if config["env-instance"] != "env-instance" {
		t.Errorf("config[\"env-instance\"] != \"env-instance\": env-instance is:%s", config["environment"])
	}
	if config["local"] != "local" {
		t.Errorf("config[\"local\"] != \"default\": local is:%s", config["local"])
	}

	os.Setenv("GO_ENV", "")
	os.Setenv("GO_APP_INSTANCE", "")
	os.Setenv("GO_PROFILES_ACTIVE", "")
	config = LoadConfigWithDir("../test/config/yaml")
	if config["local"] != "local" {
		t.Errorf("config[\"local\"] != \"local\": local is:%s", config["local"])
	}
	if config["local-development"] != "local-development" {
		t.Errorf("config[\"local-development\"] != \"local-development\": local-development is:%s", config["local-development"])
	}
}

func TestPropertiesConfigIsOverloaded(t *testing.T) {

	os.Setenv("GO_ENV", "environment")
	os.Setenv("GO_APP_INSTANCE", "instance")
	os.Setenv("GO_PROFILES_ACTIVE", "1,2")
	var config map[string]interface{}
	config = LoadConfigWithDir("../test/config/properties")

	if config["default"] != "default" {
		t.Errorf("config[\"default\"] != \"default\": default is:%s", config["default"])
	}
	if config["application"] != "application" {
		t.Errorf("config[\"application\"] != \"application\": application is:%s", config["application"])
	}
	if config["environment"] != "environment" {
		t.Errorf("config[\"environment\"] != \"environment\": environment is:%s", config["environment"])
	}
	if config["env-instance"] != "env-instance" {
		t.Errorf("config[\"env-instance\"] != \"env-instance\": env-instance is:%s", config["environment"])
	}
	if config["local"] != "local" {
		t.Errorf("config[\"local\"] != \"default\": local is:%s", config["local"])
	}

	os.Setenv("GO_ENV", "")
	os.Setenv("GO_APP_INSTANCE", "")
	os.Setenv("GO_PROFILES_ACTIVE", "")
	config = LoadConfigWithDir("../test/config/properties")
	if config["local"] != "local" {
		t.Errorf("config[\"local\"] != \"local\": local is:%s", config["local"])
	}
	if config["local-development"] != "local-development" {
		t.Errorf("config[\"local-development\"] != \"local-development\": local-development is:%s", config["local-development"])
	}
}

func TestNestedPropertiesConfigIsOverloaded(t *testing.T) {
	var config map[string]interface{}
	config = LoadConfigWithDir("../test/config/nesting")

	level1 := config["level1"].(map[string]interface{})
	level2 := level1["level2"].(map[string]interface{})
	list := level2["list"].([]interface{})
	list2 := level2["list2"].([]interface{})

	if level2["key1"].(string) != "default" {
		t.Errorf("level2[\"key1\"].(string) != \"default\"; \"default\" is:%s", level2["key1"].(string))
	}
	if level2["key2"].(string) != "application" {
		t.Errorf("level2[\"key2\"].(string) != \"default\";  \"application\" is:%s", level2["key2"].(string))
	}
	if list[0] != "application" {
		t.Errorf("list[0] != \"application\"; application is %s", list[0].(string))
	}
	if list[1] != "application" {
		t.Errorf("list[1] != \"application\"; application is %s", list[1].(string))
	}
	key3 := list[2].(map[string]interface{})["default"].(map[string]interface{})["key3"].(string)
	if key3 != "default" {
		t.Errorf("list[2].default.key3 is %s", list[1].(string))
	}
	if list2[0] != "application" {
		t.Errorf("list2[0] != \"application\"; application is %s", list[1].(string))
	}
	if list2[1] != "application" {
		t.Errorf("list2[1] != \"application\"; application is %s", list[1].(string))
	}
}

//func TestPropertiesParse(t *testing.T) {
//	os.Setenv("GO_ENV", "environment")
//	os.Setenv("GO_APP_INSTANCE", "instance")
//	os.Setenv("GO_PROFILES_ACTIVE", "1,2")
//	var config map[string]interface{}
//	config = LoadConfigWithDir("../test/config/properties")
//	t.Log(config)
//}
