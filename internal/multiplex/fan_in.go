package multiplex

import "sync"

func FanIn[T any](in ...<-chan T) <-chan T {
	c := make(chan T)
	wg := &sync.WaitGroup{}
	for i := range in {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for obj := range in[i] {
				c <- obj
			}
		}()
	}
	go func() {
		defer close(c)
		wg.Wait()
	}()
	return c
}
