package main

import (
  "fmt"
)
type Animal struct {

  name string
  sound string
  legs int
}

func modifier(animal *Animal){
  // testing pointer manipulation
  (*animal).legs=90
}

func main(){

  fmt.Println("testing struct")
  dog := Animal{"Bruno", "bhow", 4}

  fmt.Println("animal is:", dog)
  modifier(&dog)
  fmt.Println("animal is:", dog)
}
