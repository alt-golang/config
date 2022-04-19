package config

import "reflect"

type ResolverCallBack func(interface{}) interface{}

type Resolver interface {
	Resolve(object interface{}) (interface{}, error)
	CallBack(interface{}) interface{}
}

func ResolverMapValuesDeep(values interface{}, callback ResolverCallBack) interface{} {
	if reflect.TypeOf(values).Kind() == reflect.Map {
		result := map[string]interface{}{}
		for key, value := range values.(map[string]interface{}) {
			result[key] = ResolverMapValuesDeep(value, callback)
		}
		return result
	}

	return callback(values)
}
