package main

import (
	"fmt"
	"github.com/goldingp/go-talk-concurrency/internal/lessboring"
)

func main() {
	fmt.Println("endlessly ...")
	lessboring.Boring("less boring!")
}
