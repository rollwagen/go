package learngowithtests

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Markus")
	want := "Hello, Markus!"
	if got != want {
		t.Errorf("received: %q, but wanted %q", got, want)
	}
}
