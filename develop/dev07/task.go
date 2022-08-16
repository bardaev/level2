package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{} = func(channels ...<-chan interface{}) <-chan interface{} {
		out := make(chan interface{})
		var wg sync.WaitGroup
		wg.Add(len(channels))
		for _, c := range channels {
			go func(ch <-chan interface{}) {
				for v := range ch {
					out <- v
				}
				wg.Done()
			}(c)
		}
		go func() {
			wg.Wait()
			close(out)
		}()
		return out
	}
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),

		// sig(2*time.Second),
		// sig(5*time.Second),
		// sig(1*time.Second),
		// sig(1*time.Second),
		// sig(1*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
