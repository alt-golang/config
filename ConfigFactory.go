package config

import (
	"github.com/alt-golang/config/loader"
	"os"
)

func GetConfig() Config {
	return ConfigFactory(loader.LoadConfig())
}

func GetConfigWithDir(dir string) Config {
	return ConfigFactory(loader.LoadConfigWithDir(dir))
}

func ConfigFactory(config map[string]interface{}) Config {
	passphrase := os.Getenv("GO_CONFIG_PASSPHRASE")
	if passphrase == "" {
		passphrase = "changepassphrase"
	}

	var placeHolderResolver = PlaceHolderResolver{
		selector: PlaceHolderSelector{},
	}

	var gosyptDecryptor = GosyptDecryptor{
		selector: PrefixSelector{
			prefix: "enc.",
		},
		passphrase: passphrase,
	}

	var resolvers = []Resolver{nil, nil}
	var delegatingResolver = DelegatingResolver{
		resolvers: resolvers,
	}

	var defaultConfig = DefaultConfig{
		object: config,
		path:   "",
	}

	var valueResolvingConfig = ValueResolvingConfig{
		delegatingConfig: DelegatingConfig{
			defaultConfig: defaultConfig,
		},
		resolver: delegatingResolver,
	}

	placeHolderResolver.config = valueResolvingConfig
	resolvers[0] = placeHolderResolver
	resolvers[1] = gosyptDecryptor

	return valueResolvingConfig
}

var config = GetConfig()

func Has(path string) bool {
	return config.Has(path)
}
func Get(path string) (interface{}, error) {
	return config.Get(path)
}
func GetWithDefault(path string, defaultValue interface{}) (interface{}, error) {
	return config.GetWithDefault(path, defaultValue)
}
