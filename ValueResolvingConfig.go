package config

type ValueResolvingConfig struct {
	delegatingConfig DelegatingConfig
	resolver         Resolver
	Config
}

func (valueResolvingConfig ValueResolvingConfig) Get(path string) (interface{}, error) {
	object, err := valueResolvingConfig.delegatingConfig.Get(path)
	if err != nil {
		return nil, err
	}
	return valueResolvingConfig.resolver.Resolve(object)
}

func (valueResolvingConfig ValueResolvingConfig) GetWithDefault(path string, defaultValue interface{}) (interface{}, error) {
	object, err := valueResolvingConfig.delegatingConfig.GetWithDefault(path, defaultValue)
	if err != nil {
		return nil, err
	}
	return valueResolvingConfig.resolver.Resolve(object)
}

func (valueResolvingConfig ValueResolvingConfig) Has(path string) bool {
	return valueResolvingConfig.delegatingConfig.Has(path)
}
