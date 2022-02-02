package main

import "testing"

func TestToTest(t *testing.T) {
	want := "Hello"
	if got := ToTest(); got != want {
		t.Errorf("ToTest() = %s, want %s", got, want)
	}
}
