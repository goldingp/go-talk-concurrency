package lessboring

import (
	"fmt"
	"math/rand"
	"time"
)

func Boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
