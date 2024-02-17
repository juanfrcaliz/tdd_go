package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Cata")
		want := "Hello, Cata"

		CheckTest(t, got, want)
	})
	t.Run("say 'Hello, World' when an emtpy string is supplied", func(t *testing.T){
		got := Hello("")
		want := "Hello, World"

		CheckTest(t, got, want)
	})
}

func CheckTest(t testing.TB, got string, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
