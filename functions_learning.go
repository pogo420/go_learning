// entry point is must. package main is entry point in go. 
// In main package we need to have function main which returns no value.
// passing a pointer is faster and memory efficient
// maps and slices act as pass by reference 
// returning a variable as pointer from a function ? go promotes the variable in to head from stack memory.
// returning multiple values by comma only return a, b
// nil in place of None or null
// we have method in go, its defined with the context of struct. IMP we get pass by value instance.

package main

import (
	"fmt"
)

type greeter struct {
	name string
	greetings string
}


func main(){

val := test("ola")
fmt.Println(val)

f := func() {
fmt.Println("Go anonymous")
	}

f()

gr := greeter { 
	name: "French",
	greetings: "OLA",
}
gr.greet()

}


func test(name string) string{

	return name+"---"

}

func (g greeter) greet(){
	fmt.Println(g.name, g.greetings)
}
