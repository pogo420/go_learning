// pointers stores the address of the types variable
// & is called aaddress off operator
// * is for dereference
// IMP: pointer arithmetic is not allowed in go directly , use unsafe package for these operations.
// null pointers are assigned value <nil> unlike c or cpp
// new() is used for creating empty object new(struct) returns object of struct
// slice and maps does not act as deep copy(REM)

package main

import ( "fmt")

func main(){
	var a int = 42
	var a_ptr *int  // delaire a pointer
	a_ptr = &a  // & returns the address
	fmt.Println(a,a_ptr, *a_ptr, &a_ptr)  // * returns the value of that address
	*a_ptr ++ // increasing the value 
	fmt.Println(a,a_ptr, *a_ptr, &a_ptr)  // * returns the value of that address
}
