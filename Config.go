package config

type Config interface {
	Has(path string) bool
	Get(path string) (interface{}, error)
	GetWithDefault(path string, defaultValue interface{}) (interface{}, error)
}
