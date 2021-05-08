package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Brendan")

	got := buffer.String()
	want := "Hello, Brendan"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
