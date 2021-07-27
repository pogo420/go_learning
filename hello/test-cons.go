package main

import "fmt"

const VAR string = "KOLO))("
const NUM int = 34

func Test_defer() {
	fmt.Println("something in test_defer")
	defer fmt.Println("cleanup test-defer")
}

/*
func main() {

	fmt.Println(VAR,NUM)
	defer fmt.Println("main cleanup")
	test_defer()
}
*/
