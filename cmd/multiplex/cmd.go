package main

import (
	"fmt"
	"github.com/goldingp/go-talk-concurrency/internal/multiplex"
	"github.com/google/uuid"
	"strings"
	"time"
)

func main() {
	fmt.Println("Will multiplex any number of less boring services and print their messages without sequence. The whole command will timeout after eight seconds unless all services finish creating messages by then.")

	var c <-chan multiplex.MessageCounter[string]
	c = multiplex.FanIn(
		multiplex.Boring("Joe"),
		multiplex.Boring("Ann"),
		multiplex.Boring("Bob"),
		multiplex.Boring("Kim"))
	processMessages(c)

	fmt.Println()

	var err error
	var d <-chan multiplex.MessageCounter[uuid.UUID]
	alpha, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	beta, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	charlie, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	d = multiplex.FanIn(
		multiplex.Boring(alpha),
		multiplex.Boring(beta),
		multiplex.Boring(charlie))
	processMessages(d)
}

func processMessages[T any](c <-chan multiplex.MessageCounter[T]) {
	timeout := time.After(time.Duration(8) * time.Second)
	for {
		select {
		case <-timeout:
			fmt.Println("You're too slow!")
			return
		case msg, open := <-c:
			switch open {
			case false:
				fmt.Println("No more messages.")
				return
			case true:
				b := new(strings.Builder)
				b.WriteString(fmt.Sprintf("%v %d", msg.Message, msg.Count))
				if msg.Last {
					b.WriteString("\tThis less boring service has finished.")
				}
				fmt.Println(b.String())
			}
		}
	}
}
