package config

type DefaultResolver struct {
	Resolver
}

func (defaultResolver DefaultResolver) Callback(value interface{}) interface{} {
	return value
}
func (defaultResolver DefaultResolver) Resolve(object interface{}) (interface{}, error) {
	return object, nil
}
