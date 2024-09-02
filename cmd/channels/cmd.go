package main

import (
	"fmt"
	"github.com/goldingp/go-talk-concurrency/internal/channels"
)

func main() {
	fmt.Println("Prints five less boring messages using channels.")
	c := make(chan string)
	go channels.Boring("less boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You said: %q\n", <-c)
	}
	fmt.Println("You're boring! I'm leaving.")
}
