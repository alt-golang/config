package config

import "strings"

type PlaceHolderSelector struct {
	Selector
}

func (placeHolderSelector PlaceHolderSelector) Matches(value string) bool {
	return strings.Contains(value, "${") && strings.Contains(value, "}") &&
		strings.Index(value, "${") < strings.Index(value, "}")
}

func (placeHolderSelector PlaceHolderSelector) ResolveValue(value string) string {
	return value
}
