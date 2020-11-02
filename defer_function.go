// all defer in LILO
// before returning a function: defer statement are executed.
// defer is perfect for cleanup operations.
// defer is evaluated based on the values at the time when, defer is called not during execution.
package main

import (
	"fmt"
)

func main(){

	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}
