
package main

import "fmt"

func main(){

  fmt.Println("Testing pointers..")

  var data int = 789
  var data_ptr *int = &data

  var arr= [3]int{12,45,8}

  var arr_ptr *[3]int = &arr

  fmt.Println(data, data_ptr, *data_ptr)
  *data_ptr = 890
  fmt.Println(data, data_ptr, *data_ptr)
  fmt.Println(arr, arr_ptr, *arr_ptr)
  *arr_ptr = [3]int{45,897,0}
  fmt.Println(arr, arr_ptr, *arr_ptr)
}

