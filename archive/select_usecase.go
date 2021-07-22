// A select blocks until one of its cases can run,
// then it executes that case. It chooses one at random if multiple are ready.

package main
import "fmt"


func sample_conc(c, quit chan int){
  x:=0
  for {
    select {
    case c <-x : // once this is done, it is moved to next case
      x++
    case <- quit:
      fmt.Println("Quit")
      return
    }
  }
}

func main(){
  c := make(chan int)
  quit := make(chan int)

  go func ()  {
    for i:=0; i < 10; i++{
      fmt.Println(<-c)
    }
    //close(c)
    quit <- 0
  } ()
  sample_conc(c, quit)
}
