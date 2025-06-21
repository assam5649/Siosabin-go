package test

import "testing"

func Failure() string {
	return "Hello, World!"
}

func TestFailure(t *testing.T) {
	got := HelloWorld()
	want := "Goodbye, World!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
