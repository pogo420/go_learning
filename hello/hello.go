package main

import "fmt"
import "rsc.io/quote"
import "greetings"
import "log"

func main() {

    log.SetPrefix("sample-logs: ")
    log.SetFlags(0)

    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())
    message, error := greetings.Hello("")
    if error != nil {
        log.Fatal("Empty string")  // prints and terminates the program
	}
	fmt.Println(message)
}

