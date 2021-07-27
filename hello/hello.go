package main

import (
	"fmt"
	"greetings"
	"log"
	"rsc.io/quote"
  //"advancetypes/advancetypes"
  //"main/testcons"
)

func main() {

	log.SetPrefix("sample-logs: ")
	log.SetFlags(0)
  //advancetypes.Ptr()
  //testcons.Test_defer()
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	message, error := greetings.Hello("")
	if error != nil {
		log.Fatal("Empty string") // prints and terminates the program
	}
	fmt.Println(message)

}
