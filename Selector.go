package config

type Selector interface {
	Matches(value string) bool
	ResolveValue(value string) string
}
