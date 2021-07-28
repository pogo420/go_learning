package main

import (
  "fmt"
)

func main(){
  var arr [5]int = [5]int{1,2,3,4,5}

  var s[]int = arr[1:3]
  fmt.Println(arr, s)

  var s1[]int
  fmt.Println(s1)
  for i:=0; i< 10; i++ {
    s1=append(s1,i)
    fmt.Println(s1)
  }

fmt.Println("go range check..")

for index, value := range arr {
    fmt.Println(index, value)
  } 
}
