package main

import "testing"

func TestSayHello(t *testing.T) {
	hello := SayHello()

	if hello != "Hello World!" {
		t.Errorf("Expected %s to equal %s", hello, "Hello World!")
	}
}
