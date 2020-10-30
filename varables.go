// showing cleaner varable delaration
// lower case variables are scoped to the packages.
// upper case variables are scoped outside the package as well.
// casting is similar int() float32()
// all varables must be used

package main
import (
	"fmt"
	"strconv"
)

var (
	values int = 32
	name string = "gupai"
)

var GLOB string = "loko" // global

func main(){

	fmt.Printf("%v, %T", values, values)
	fmt.Println("")
	fmt.Printf("%v, %T", name, name)
	fmt.Println(strconv.Itoa(values)) // string conversion

}
