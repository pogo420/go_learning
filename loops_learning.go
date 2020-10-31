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

}
