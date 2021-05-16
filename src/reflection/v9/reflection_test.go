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
		{
			"Pointers to things",
			&Person{
				"Brendan",
				Profile{24, "China"},
			},
			[]string{"Brendan", "China"},
		},
		{
			"Slices",
			[]Profile{
				{24, "China"},
				{33, "London"},
			},
			[]string{"China", "London"},
		},
		{
			"Arrays",
			[2]Profile{
				{24, "China"},
				{33, "London"},
			},
			[]string{"China", "London"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{24, "China"}
			aChannel <- Profile{33, "Berlin"}
			close(aChannel)
		}()

		var got []string
		want := []string{"China", "Berlin"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
