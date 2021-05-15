package main

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Brendan"},
			[]string{"Brendan"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Brendan", "China"},
			[]string{"Brendan", "China"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Brendan", 24},
			[]string{"Brendan"},
		},
		{
			"Nested fields",
			Person{
				"Brendan",
				Profile{24, "China"},
			},
			[]string{"Brendan", "China"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				// 先处理第一个参数x-test.Input，待walk处理到fn匿名函数时，再执行下方语句
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
