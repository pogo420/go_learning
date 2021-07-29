package main

import (
  "fmt"
)

type Person struct {
  name string
  id int
}

// method example 1
func (v Person)getName()string{
  return v.name
}

// method example 2
func (v Person)getId()int{
  return v.id
}

func main(){
  fmt.Println("testing methods")
  pr := Person{"gupeh", 19}
  fmt.Println(pr.getName(), pr.getId())
  var arr [4]int = [4]int{1,3,45,8}
  fmt.Println(arr)
}
