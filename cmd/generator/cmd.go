package main

import (
	"fmt"
	"github.com/goldingp/go-talk-concurrency/internal/generator"
)

func main() {
	fmt.Println("Prints five less boring messages using channels and the generator pattern.")
	c := generator.Boring("less boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You said %q\n", <-c)
	}
	fmt.Println("You're boring! I'm leaving.")
}
