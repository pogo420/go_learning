//flow of execution main/function --> defer --> panic
//

package main

import (
	"fmt"
)

func main(){

	fmt.Println("Ola")
	panic("SOmething happned")
	fmt.Println("OlaOla")

}
