package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Cata")
	want := "Hello, Cata"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
