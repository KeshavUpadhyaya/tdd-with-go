package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Keshav")

	got := buffer.String()
	want := "Hello, Keshav"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
