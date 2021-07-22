//v, ok := <-ch
// ok is false if there are no more values to receive and the channel is closed.
// The loop for i := range c receives values from the channel repeatedly until it is closed.
// sender can close the channel, never the receiver.
// Sending on a closed channel will cause a panic.
// Channels aren't like files;
// you don't usually need to close them. Closing is only necessary when the
// receiver must be told there are no more values coming, such as to terminate a range loop.

package main

import (
  "fmt"
  "math/rand"
)

func generateRand(limit int, channel *chan int){
  var i int = 0
  for (i < limit){
      *channel <- rand.Intn(100)
      i++
  }
  close(*channel) // if we dont give close, it will create a panic.
}

func main(){
  channel := make(chan int)
  go generateRand(100, &channel)
  for i := range channel {
    fmt.Println(i)
  }
}
