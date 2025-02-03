package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("chirs")
	want := "Hello, chirs"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
