package main

import "testing"

func TestAdd(t *testing.T) {
	x := 1
	y := 4
	result := Add(x, y)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}
