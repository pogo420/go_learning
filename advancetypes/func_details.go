package main

import (
  "fmt"
  s "strings"
)

func process_string(value string, function func(string)string){
  fmt.Println(function(value))
}

func main(){

  // function definition 
  uc_func := func(value string)string{
    return s.ToUpper(value)
  }
  process_string("hello", uc_func)
}
