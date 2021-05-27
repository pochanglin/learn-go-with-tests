package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkVal := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkVal(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkVal(val.Index(i))
		}
	case reflect.String:
		fn(val.String())
	case reflect.Map:
		for _, k := range val.MapKeys() {
			walk(val.MapIndex(k).Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
