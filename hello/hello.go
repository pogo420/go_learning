package main

import "fmt"
import "rsc.io/quote"
import "greetings"
import "log"
import "math/rand"
import "time"


func main() {

    log.SetPrefix("sample-logs: ")
    log.SetFlags(0)

    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())
    /*
    message, error := greetings.Hello("")
    if error != nil {
        log.Fatal("Empty string")  // prints and terminates the program
	}
	fmt.Println(message)
	*/

	fmt.Println(randomFormat())

	
}

func randomFormat() string {
    // Array
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}

// special purpose function: executed before main()
// used for setting global variables
// init sets initial values for variables used in the function.
// no input or no output.
// whenever a package is imported init is executed.
// it is avalable per package. 
func init() {
    rand.Seed(time.Now().UnixNano())
}

func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages. Memory allocation 
    messages := make(map[string]string)
    
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {

        message, err := Hello(name)

        if err != nil {
            return nil, err
        }

        // In the map, associate the retrieved message with
        // the name.

        messages[name] = message
    }
    return messages, nil
}

