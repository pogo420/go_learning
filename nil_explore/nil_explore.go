package main

import "fmt"

func isError(err error) bool {
	if err != nil {
		return true
	} else {
		return false
	}
}

func something() error {
	return nil
}

func main() {

	if something() == nil {
		fmt.Println("nil equals")
	}

	if !isError(something()) {
		fmt.Println("nil equals-function")
	}

}
