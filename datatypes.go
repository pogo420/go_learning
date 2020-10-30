package main
// defalt value for varables are zero
// operations between mixed data types are not allowed in go.
import "fmt"

var flag bool

func main(){
	flag = true
	if flag{
	fmt.Println("Ola")
	}
}
