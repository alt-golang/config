package config

import (
	gosypt "github.com/alt-golang/gosypt.pkg"
	"reflect"
)

type GosyptDecryptor struct {
	selector   Selector
	passphrase string
	Resolver
}

func (gosyptDecryptor GosyptDecryptor) Callback(value interface{}) interface{} {
	if reflect.TypeOf(value).Kind() == reflect.String &&
		gosyptDecryptor.selector.Matches(value.(string)) {
		selectedValue := gosyptDecryptor.selector.ResolveValue(value.(string))
		result, err := gosypt.DecryptString(gosyptDecryptor.passphrase, selectedValue)
		if err == nil {
			return result
		}
	}
	return value
}

func (gosyptDecryptor GosyptDecryptor) Resolve(object interface{}) (interface{}, error) {
	return ResolverMapValuesDeep(object, gosyptDecryptor.Callback), nil
}
