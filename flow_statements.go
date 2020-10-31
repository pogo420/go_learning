// No need for break, fallthrough is not compulsory
// switch can be used for type check as well  
package main

import (
	"fmt"
)

func main(){
var i int = 4

switch(i){
case 1,5,10:
	fmt.Println("Ola1")
case 3:
	fmt.Println("Ola2")
default:
	fmt.Println("default")
}
}

