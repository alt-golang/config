package config

type DelegatingConfig struct {
	defaultConfig DefaultConfig
	Config
}

func (config DelegatingConfig) Get(path string) (interface{}, error) {
	return config.defaultConfig.Get(path)
}

func (config DelegatingConfig) GetWithDefault(path string, defaultValue interface{}) (interface{}, error) {
	return config.defaultConfig.GetWithDefault(path, defaultValue)
}

func (config DelegatingConfig) Has(path string) bool {
	return config.defaultConfig.Has(path)
}
