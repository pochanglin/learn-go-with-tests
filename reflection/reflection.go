package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Struct:
		numOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice:
		numOfValues = val.Len()
		getField = val.Index
	case reflect.String:
		fn(val.String())
	}

	for i := 0; i < numOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
