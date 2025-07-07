package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello World"
	got := Hello()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
