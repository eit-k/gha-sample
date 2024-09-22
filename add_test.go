package main

import "testing"

func TestAdd(t *testing.T) {
	x := 2
	y := 3
	result := Add(x, y)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}
