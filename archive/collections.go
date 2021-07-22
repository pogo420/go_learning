// script for arrays and slices --> collection in go.
// assigning array to another variable will copy it. (inside function) deep copy.
// slice and array are similar but slice is like python list.

package main

import(
	"fmt"
)

func main() {

	// declaring array can hold 3 elements fo type int
	grades:=[3]int {90, 89, 78} // array

	fmt.Println(grades)
	fmt.Println(grades[2]) // indexing 
	
	var stud [3]string
	fmt.Println(stud)
	stud[2] = "gupau"
	fmt.Println(stud)
	fmt.Println(len(stud))
	
	slc:= make([]int, 1,10)  // creates a slice
	//slc:=[]int{1,2,3} // slice 
	fmt.Println(slc, len(slc), cap(slc))	
	
	slc = append(slc, 10)
	fmt.Println(slc, len(slc), cap(slc))	
	slc = append(slc, []int{12,89,70}...)  // spreading operator
	fmt.Println(slc, len(slc), cap(slc))	

}
