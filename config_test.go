package config

import (
	"testing"
)

func TestConfig(t *testing.T) {

	config := GetConfig()
	if config.Has("default") == false {
		t.Errorf("config.Has(\"default\"): config.Has(\"default\") is:%t", config.Has("default"))
	}
	if config.Has("nokey") == true {
		t.Errorf("config.Has(\"nokey\") == false: config.Has(\"nokey\") is:%t", config.Has("nokey"))
	}
	def, _ := config.Get("default")
	if def.(string) != "default" {
		t.Errorf("default !=\"default\": default is:%s", def.(string))
	}
	nokey, _ := config.GetWithDefault("nokey", "default")
	if nokey.(string) != "default" {
		t.Errorf("nokey !=\"default\": nokey is:%s", def.(string))
	}

	val, _ := config.Get("placeholder")
	if val.(string) != "start.one.two.end" {
		t.Errorf("config.Get(\"placeholder\") != \"start.one.two.end\": placeholder is:%s", val.(string))
	}
	val, _ = config.Get("nested.encrypted")
	if val.(string) != "hello world" {
		t.Errorf("config.Get(\"nested.encrypted\") != \"hello world\": nested.encrypted is:%s", val.(string))
	}
}
