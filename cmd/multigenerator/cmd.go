package main

import (
	"fmt"
	"github.com/goldingp/go-talk-concurrency/internal/generator"
)

func main() {
	fmt.Println("Prints five messages from each of two less boring services.")
	joe := generator.Boring("Joe")
	ann := generator.Boring("Ann")

	for i := 0; i < 5; i++ {
		fmt.Printf("%s\n", <-joe)
		fmt.Printf("%s\n", <-ann)
	}

	fmt.Println("You're both boring! I'm leaving.")
}
