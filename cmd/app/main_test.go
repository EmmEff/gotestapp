package main

import "testing"

func TestBool(t *testing.T) {
	if !true {
		t.Fatal("What is going on here?!?!")
	}
}
