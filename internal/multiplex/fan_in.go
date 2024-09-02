package multiplex

func FanIn[T any](in ...<-chan T) <-chan T {
	c := make(chan T)
	closed := make(chan bool, len(in))
	for i := range in {
		go func() {
			for obj := range in[i] {
				c <- obj
			}
			closed <- true
		}()
	}
	go func() {
		for i := 0; i < len(in); i++ {
			<-closed
		}
		close(c)
		close(closed)
	}()
	return c
}
