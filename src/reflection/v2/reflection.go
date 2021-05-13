package main

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	fmt.Println(val)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Println(field)
		fn(field.String())
	}
}
