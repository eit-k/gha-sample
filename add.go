package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func main() {
	x := 1
	y := 2
	result := Add(x, y)
	fmt.Println("Result of addition: ", result)
}
