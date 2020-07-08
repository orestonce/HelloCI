package main

import "testing"

func TestGetMessage(t *testing.T) {
	if "Hello world" != GetMessage() {
		t.Fatal("GetMessage invalid")
	}
}
