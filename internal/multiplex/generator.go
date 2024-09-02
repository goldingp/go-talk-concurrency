package multiplex

import (
	"math/rand"
	"time"
)

func Boring[T any](msg T) <-chan MessageCounter[T] {
	c := make(chan MessageCounter[T])
	go func() {
		timeout := time.After(time.Duration(rand.Intn(8e3)) * time.Millisecond)
		for i := 0; ; i++ {
			select {
			case <-timeout:
				c <- MessageCounter[T]{Count: i, Last: true, Message: msg}
				close(c)
				return
			case c <- MessageCounter[T]{Count: i, Last: false, Message: msg}:
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
		}
	}()
	return c
}
