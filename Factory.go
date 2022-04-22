package config

import (
	"os"
)

func Factory(config map[string]interface{}) Config {
	passphrase := os.Getenv("GO_CONFIG_PASSPHRASE")
	if passphrase == "" {
		passphrase = "changepassphrase"
	}

	var placeHolderResolver = PlaceHolderResolver{
		selector: PlaceHolderSelector{},
	}

	var urlResolver = URLResolver{
		selector: PrefixSelector{
			prefix: "url.",
		},
	}

	var gosyptDecryptor = GosyptDecryptor{
		selector: PrefixSelector{
			prefix: "enc.",
		},
		passphrase: passphrase,
	}

	var resolvers = []Resolver{nil, nil, nil}
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
	urlResolver.config = valueResolvingConfig

	resolvers[0] = placeHolderResolver
	resolvers[1] = urlResolver
	resolvers[2] = gosyptDecryptor

	return valueResolvingConfig
}
