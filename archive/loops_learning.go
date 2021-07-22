// for is the only keyword for loop
// for can be used as while 


package main

import (
	"fmt"
)


func main(){

	for i:=0; i< 10; i++{	
	fmt.Println(i)
	}

i:=0
//while style
	for i< 10 {	
	fmt.Println(i)
	i++
	}
// do while style 
for{
i++
fmt.Println(i)
if i > 20{
break
}
}
arr := [5]int{34,5,6,789,9}
fmt.Println("===============================")
for i, v:= range arr{ // iterator vs values
fmt.Println(i, v)
}

}
