package main

import (
	"fmt"
)

func add(x int) (int, int) {
	return x + 1, x + 2
}

func main() {
	x, y := add(45) // unpacking values
	fmt.Println(x, y)
}
