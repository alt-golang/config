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

	t.Log(config)
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
	t.Log(config)
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
	t.Log(config)
	t.Log(config["level1"])
	t.Log(config["level1"].(map[string]interface{})["level2"])

	level1 := config["level1"].(map[string]interface{})
	level2 := level1["level2"].(map[string]interface{})

	if level2["key1"].(string) != "default" {
		t.Errorf("level2[\"key1\"].(string) != \"default\" != \"default\" is:%s", level2["key1"].(string))
	}
	if level2["key2"].(string) != "application" {
		t.Errorf("level2[\"key2\"].(string) != \"default\" != \"application\" is:%s", level2["key2"].(string))
	}
	t.Log("done")
}

//func TestPropertiesParse(t *testing.T) {
//	os.Setenv("GO_ENV", "environment")
//	os.Setenv("GO_APP_INSTANCE", "instance")
//	os.Setenv("GO_PROFILES_ACTIVE", "1,2")
//	var config map[string]interface{}
//	config = LoadConfigWithDir("../test/config/properties")
//	t.Log(config)
//}
