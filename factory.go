package config

import (
	"alt-golang/config/loader"
	"os"
)

func GetConfig() Config {

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
		object: loader.LoadConfig(),
		path:   "",
	}

	var config = ValueResolvingConfig{
		delegatingConfig: DelegatingConfig{
			defaultConfig: defaultConfig,
		},
		resolver: delegatingResolver,
	}

	placeHolderResolver.config = config
	resolvers[0] = placeHolderResolver
	resolvers[1] = gosyptDecryptor

	return config
}

//var config = GetConfig()
