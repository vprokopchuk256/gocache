package gcpipe_test

import (
	"fmt"
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcpipe"
)

func TestMultiplexer(t *testing.T) {
	in1 := make(chan error)
	defer close(in1)

	in2 := make(chan error)
	defer close(in2)

	done := make(chan interface{})
	defer close(done)

	out := gcpipe.Multiplexer(done, in1, in2)

	go func() {
		in1 <- fmt.Errorf("err1")
		in2 <- fmt.Errorf("err1")

		done <- true
	}()

	<-done
	err1 := <-out
	err2 := <-out

	if err1 == nil || err2 == nil {
		fmt.Errorf("expected input channels to be multiplexed")
	}
}
