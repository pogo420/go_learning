package main

import (
  "fmt"
)

type Person struct {
  name string
  id int
}


func main(){
  fmt.Println("Map Testing")
  // make is cumpulsory assigning value 
  var person_list map[int]Person = make(map[int]Person)
  fmt.Println(person_list)
  person_list[12] = Person{"Arnab", 12}
  fmt.Println(person_list)
}
