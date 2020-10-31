// main is must in go like java
// struct should end with comma
// Exact inheritance is not present in go, we have composition(Works cool)

package main

import (
	"fmt"
)

type Doctor struct {
	id int
	name string
	grades [5]int
}

func main(){
	adoc := Doctor{
		id: 90,
		name: "pupau",
		grades: [5]int {32, 42, 45, 35, 40},
	}

 fmt.Println(adoc)
 fmt.Println(adoc.grades)

// anonymous struct 
anStruct := struct{name string}{name: "Ola"}
fmt.Println(anStruct)

// showing composition

type Animal struct {
	Name string
	Origin string
}

type Bird struct {
	Animal
	speed float32
	CanFly bool
}

b := Bird{}
b.Name="Oa"
b.Origin="India"
b.speed=89.78
b.CanFly=true

fmt.Println(b)

}

