package main

import (
  "fmt"
  s "strings"
  "time"
)


func api_call_dummy(data string, channel chan string){
    // function to imitate API calls
    time.Sleep(time.Duration(5*time.Second))
    channel <- s.ToUpper(data)
}

func main(){
  var string_list[5]string = [5]string{"hello godzilla", "gupai is great", "90ioki loho", "boo booo", "winter is there"}
  var result_list[5] string
  channel := make(chan string)

  for _, value:= range string_list{
    // concurrent api calls
    go api_call_dummy(value, channel)
  }

  for i, _:= range result_list{
    // collecting results
      temp:= <-channel
      result_list[i]=temp
  }
  fmt.Println(result_list)
}
