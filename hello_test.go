package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Isaac")
	want := "Hello, Isaac"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
