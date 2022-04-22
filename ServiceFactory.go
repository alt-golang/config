package config

func ServiceFactory(config map[string]interface{}) Config {

	var placeHolderResolver = PlaceHolderResolver{
		selector: PlaceHolderSelector{},
	}

	var resolvers = []Resolver{nil}
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

	return valueResolvingConfig
}
