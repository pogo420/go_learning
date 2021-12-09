package main

import "fmt"

type A struct {
	name string
	id   int
}

func (a *A) Print() {
	fmt.Println("A")
}

type A_ struct {
	A
	name    string
	address string
}

func (a *A_) Print() {
	fmt.Println("A_")
}

func main() {
	ptr := A_{}

	ptr.Print()
	ptr.A.Print()
}
