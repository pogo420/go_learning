// slice can not be map index
// map should end with comma
// maps we not ordered like python dictionary.
// delete using builtin function
// for index not present it returns zero.
// for slice and maps we we assign to another varable its passed like references
package main
import (
	"fmt"
)

func main(){
	// initializtion
	sample_map := map[string]int{
	"1": 1,
	"two": 2,
	}
	
	fmt.Println(sample_map)
	// empty map
	sample_map2 := map[string]int{
	}
	sample_map2["kolo"] = 678
	fmt.Println(sample_map2)

	_, ispresent := sample_map2["gupai"] // to check if key is present or not.
	fmt.Println(ispresent)	
}


