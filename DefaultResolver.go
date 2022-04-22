package config

type DefaultResolver struct {
	Resolver
}

func (defaultResolver DefaultResolver) Callback(value interface{}, path string) interface{} {
	return value
}
func (defaultResolver DefaultResolver) Resolve(object interface{}, path string) (interface{}, error) {
	return object, nil
}
