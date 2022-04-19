package config

type DefaultSelector struct {
	object map[string]interface{}
	path   string
	Selector
}

func (defaultSelector DefaultSelector) Matches(value string) bool {
	return true
}

func (defaultSelector DefaultSelector) ResolveValue(value string) string {
	return value
}
