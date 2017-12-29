package gcpipe

import (
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func Executor(m *gcmap.Map, done <-chan interface{}, input <-chan gccommand.Command) (<-chan string, <-chan error) {
	output := make(chan string)
	errors := make(chan error)

	go func() {
		defer close(output)
		defer close(errors)

		for {
			select {
			case <-done:
				return
			case cmd, ok := <-input:
				if ok {
					log, err := cmd.Exec(m)
					if err != nil {
						errors <- err
					} else {
						output <- log
					}
				}
			}
		}
	}()

	return output, errors
}
