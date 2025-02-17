package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Keshav", "")
		want := "Hello, Keshav"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		var got = Hello("Elodie", "Spanish")
		var want = "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		var got = Hello("Keshav", "German")
		var want = "Hallo, Keshav"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // to show line number in the function call when the test fails
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
