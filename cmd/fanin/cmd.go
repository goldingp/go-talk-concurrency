package main

import (
	"fmt"
	"github.com/goldingp/go-talk-concurrency/internal/fanin"
	"github.com/goldingp/go-talk-concurrency/internal/generator"
)

func main() {
	fmt.Println("Prints 10 messages from two less boring services, but loses sequence due to the multiplex (fan in) pattern.")
	c := fanin.FanIn(
		generator.Boring("Joe"),
		generator.Boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring! I'm leaving.")
}
