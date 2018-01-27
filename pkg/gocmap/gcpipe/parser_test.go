package gcpipe_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcpipe"
)

func TestParserValidCommand(t *testing.T) {
	input := make(chan string)
	defer close(input)

	done := make(chan interface{})
	defer close(done)

	output, _ := gcpipe.Parser(done, input)

	go func() {
		input <- "ins key 10"
		done <- true
	}()

	<-done
	<-output
}

func TestParserInvalidCommand(t *testing.T) {
	input := make(chan string)
	defer close(input)

	done := make(chan interface{})
	defer close(done)

	_, errors := gcpipe.Parser(done, input)

	go func() {
		input <- "undefined key 10"
		done <- true
	}()

	<-done
	err := <-errors

	if err == nil {
		t.Fatalf("error is expected")
	}
}
