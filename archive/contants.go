// we should assign during compile time.
// shadow works for constant ( we can overwrite)
// immutable but can be shadowed.
// iota is important operator--> successive + ve interger constant : 0, 1,2,3...

package main

import (
	"fmt")

// kind of fill in excel
const (
	_ = iota
	a = 10* iota + 4
	b
	c
)

func main(){
	fmt.Println(a,b,c)
}
