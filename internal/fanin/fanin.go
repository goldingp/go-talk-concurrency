package fanin

func FanIn(in1, in2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-in1:
				c <- s
			case s := <-in2:
				c <- s
			}
		}
	}()
	return c
}
