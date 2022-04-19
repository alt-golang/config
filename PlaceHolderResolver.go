package config

import (
	"reflect"
	"strings"
)

type PlaceHolderResolver struct {
	config   Config
	selector Selector
	Resolver
}

func (placeHolderResolver PlaceHolderResolver) Callback(value interface{}) interface{} {
	if reflect.TypeOf(value).Kind() == reflect.String &&
		placeHolderResolver.selector.Matches(value.(string)) {

		resolvedValue := ""
		remainder := value.(string)
		placeholder := ""

		loop := true
		for loop == true {

			resolvedValue = resolvedValue + remainder[0:strings.Index(remainder, "${")]
			placeholder = remainder[strings.Index(remainder, "${")+2 : strings.Index(remainder, "}")]
			lookup, err := placeHolderResolver.config.Get(placeholder)
			if err != nil {
				return value
			}
			resolvedValue = resolvedValue + lookup.(string)
			remainder = remainder[strings.Index(remainder, "}")+1:]

			loop = resolvedValue == "" ||
				(strings.Contains(remainder, "${") && strings.Contains(remainder, "}") && strings.Index(remainder, "${") < strings.Index(remainder, "}"))

		}
		resolvedValue = resolvedValue + remainder
		return resolvedValue
	}
	return value
}

func (placeHolderResolver PlaceHolderResolver) Resolve(object interface{}) (interface{}, error) {
	return ResolverMapValuesDeep(object, placeHolderResolver.Callback), nil
}
