package config

import "strings"

type PrefixSelector struct {
	prefix string
	Selector
}

func (prefixSelector PrefixSelector) Matches(value string) bool {
	return strings.HasPrefix(value, prefixSelector.prefix)
}

func (prefixSelector PrefixSelector) ResolveValue(value string) string {
	return strings.Replace(value, prefixSelector.prefix, "", 1)
}
