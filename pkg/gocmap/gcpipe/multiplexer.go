package gcpipe

import "sync"

func Multiplexer(done <-chan interface{}, channels ...<-chan error) <-chan error {
	var wg sync.WaitGroup

	multiplexedStream := make(chan error)

	multiplex := func(c <-chan error) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
