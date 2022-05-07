package main

import (
	"fmt"

	"example.com/greetings"
)

// func main() {
// 	fmt.Println(quote.Go())
// }

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Liam")
	fmt.Println(message)
}
