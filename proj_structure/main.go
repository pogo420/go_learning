package main

import "fmt"
import "proj_structure/package_1"
import "proj_structure/package_2"

func main(){
  fmt.Println("initiating main..")
  package_1.Log()
  package_1.Log2()
  package_2.Log()
  package_2.Log2()
}
