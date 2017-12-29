package gcpipe

import "github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"

func Parser(done <-chan interface{}, input <-chan string) (<-chan gccommand.Command, <-chan error) {
	output := make(chan gccommand.Command)
	errors := make(chan error)

	go func() {
		defer close(output)
		defer close(errors)

		for {
			select {
			case <-done:
				return
			case sCmd, ok := <-input:
				if ok {
					c, err := gccommand.Parse(sCmd)
					if err != nil {
						errors <- err
					} else {
						output <- c
					}
				}
			}
		}
	}()

	return output, errors
}
