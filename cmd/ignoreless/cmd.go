package main

import (
	"fmt"
	"github.com/goldingp/go-talk-concurrency/internal/lessboring"
	"time"
)

func main() {
	go lessboring.Boring("less boring!")
	fmt.Println("I'm listening.")
	time.Sleep(3 * time.Second)
	fmt.Println("You're boring! I'm leaving.")
}
