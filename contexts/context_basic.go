package main

import (
	"context"
	"log"
	"strings"
	"time"
)

func apiCalls(ctx context.Context, data string, output chan string) {
	time.Sleep(time.Duration(5 * time.Second))
	output <- strings.ToUpper(data)
}

func main() {
	/*
		Two variaties of context
		1. With Done or cancelFunc:
			a. WithCancel
			b. WithDeadline
			c. WithTimeout

		2. With value
			a. WithValue // providing value
	*/

	ctx := context.Background() // We can use context.TODO() as well same thing
	ctx_derived, cancelFunc := context.WithTimeout(ctx, 6*time.Second)

	api_response := make(chan string)
	go apiCalls(ctx_derived, "ola32", api_response)

	select {
	case data := <-api_response:
		log.Println("Data received", data)

	case <-ctx_derived.Done():
		log.Println("Time exceed..")
		cancelFunc()
	}
}
