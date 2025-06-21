package test

import "testing"

func HelloWorld() string {
	return "Hello, World!"
}

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "Hello, World!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
