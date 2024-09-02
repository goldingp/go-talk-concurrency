package main

import (
	"fmt"
	"github.com/goldingp/go-talk-concurrency/internal/boring"
)

func main() {
	fmt.Println("endlessly ...")
	boring.Boring("boring!")
}
