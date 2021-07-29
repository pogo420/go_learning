package main

import (
  "fmt"
)

type Data struct{
  data_id int
  data string
}

func (d *Data) get_data()string{
  return d.data
}


func (d *Data) get_data_id()int{
  return d.data_id
}


type Data_accessor interface{
    get_data()string
    get_data_id()int
}

func main(){

  var accessor Data_accessor = &Data{1234, "Hello Testing interface"} 
  fmt.Println(accessor.get_data())
  fmt.Println(accessor.get_data_id())
}


