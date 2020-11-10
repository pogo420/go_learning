// Sends and receives block until the other side is ready.
// This allows goroutines to synchronize without explicit locks or condition variables.
// we can create buffered channel. when buffer is full receive is blocked.

package main

import (
  "fmt"
  "time"
)


func sum(data []int, channel chan int) int {
  var sum int
  for _, value := range data{
    sum += value
  }
  return sum
}

func sum_(data []int, channel chan int) {
  var sum int
  for _, value := range data{
    sum += value
  }
  channel <-sum
}


func main()  {
  channel := make(chan int, 0)
  data := []int{1,4,6,7,9,90,-89,10000000000000000,90,89,-34,90}
  t:=time.Now()
  //fmt.Println(t)
  x := sum(data[len(data)/2:], channel)
  y := sum(data[:len(data)/2], channel)
  fmt.Println(x,y,x+y)
  fmt.Println(t.Sub(time.Now()))

  t1:=time.Now()
  //fmt.Println(t)
  go sum_(data[len(data)/2:], channel)
  go sum_(data[:len(data)/2], channel)
  c1, c2 := <-channel , <-channel
  fmt.Println(c1,c2, c1+c2)
  fmt.Println(t1.Sub(time.Now()))
}
