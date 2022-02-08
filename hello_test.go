package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("received: %q, but wanted %q", got, want)
		}
	}

	t.Run("say hello to people/names", func(t *testing.T) {
		got := Hello("Markus")
		want := "Hello, Markus!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when no name is given", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}
