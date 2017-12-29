package gcpipe_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcpipe"
)

func TestExecutorValidCMD(t *testing.T) {
	input := make(chan gccommand.Command)
	defer close(input)

	done := make(chan interface{})
	defer close(done)

	m := gcmap.New()

	output, _ := gcpipe.Executor(m, done, input)

	go func() {
		ins, _ := gccommand.NewInsert("key", "10")

		input <- ins
		done <- true
	}()

	<-done
	log := <-output

	if log != "key := (integer) 10" {
		t.Fatalf("expected '' but got '%v'", log)
	}
}

func TestExecutorInvalidCMD(t *testing.T) {
	input := make(chan gccommand.Command)
	defer close(input)

	done := make(chan interface{})
	defer close(done)

	m := gcmap.New()

	_, errors := gcpipe.Executor(m, done, input)

	go func() {
		inc, _ := gccommand.NewInc("key")

		input <- inc
		done <- true
	}()

	<-done
	err := <-errors

	if err == nil {
		t.Fatalf("error expected")
	}
}
