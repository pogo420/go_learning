// Interfaces contains method declaration 
// methods implicitly implements an interface via type context
// context of method can be of any type(user defined).
// when we do interface_type.method() it find for the implementation with defined method with the context.
// we ability to embedded interfaces like struct, combining interfaces.


package main
import (
	"fmt"
)

type Writer interface {
	write(data string) (string, error)
}

type StdWriter struct {
	parameter string
}

func (st StdWriter)write(value string) (string, error){

	return st.parameter + value, nil
}

func main(){
	fmt.Println("Hellow")
	var wr Writer = StdWriter{}
	fmt.Println(wr.write("Ola"))
}


