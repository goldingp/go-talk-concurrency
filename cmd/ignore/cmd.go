package main

import (
	"fmt"
	"github.com/goldingp/go-talk-concurrency/internal/boring"
	"github.com/goldingp/go-talk-concurrency/internal/lessboring"
)

func main() {
	fmt.Println("Ignore boring and less boring functions.")
	go boring.Boring("boring!")
	go lessboring.Boring("less boring!")
}
