package config

import "reflect"

type ResolverCallBack func(interface{}, string) interface{}

type Resolver interface {
	Resolve(object interface{}, path string) (interface{}, error)
	CallBack(interface{}, string) interface{}
}

func ResolverMapValuesDeep(values interface{}, path string, callback ResolverCallBack) interface{} {
	if reflect.TypeOf(values).Kind() == reflect.Map {
		result := map[string]interface{}{}
		for key, value := range values.(map[string]interface{}) {
			result[key] = ResolverMapValuesDeep(value, path, callback)
		}
		return result
	}

	return callback(values, path)
}
