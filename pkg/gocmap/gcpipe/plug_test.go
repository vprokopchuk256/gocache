package gcpipe_test

import (
	"fmt"
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcpipe"
)

type socket struct {
	write   string
	read    string
	err     error
	readErr error
	open    bool
}

func (s *socket) Read() (string, error) {
	if s.readErr != nil {
		return "", s.readErr
	}

	return s.read, nil
}

func (s *socket) Write(w string) error {
	s.write = w

	return nil
}

func (s *socket) Error(err error) error {
	s.err = err

	return nil
}

func (s *socket) Close() {
	s.open = false
}

func TestPlugSuccess(t *testing.T) {
	socketOutStr := "socket out str"
	plugInStr := "plug in str"
	plugInChan := make(chan string)

	s := &socket{read: socketOutStr}

	p := gcpipe.Plug(s)
	p.SetInput(plugInChan)
	p.Start()

	plugOutStr := <-p.Output()

	plugInChan <- plugInStr

	p.Close()

	if plugOutStr != socketOutStr {
		t.Errorf("expected '%v', got: '%v'", socketOutStr, plugOutStr)
	}

	if s.write != plugInStr {
		t.Errorf("expected '%v', got: '%v'", plugInStr, s.write)
	}
}

func TestPlugError(t *testing.T) {
	socketOutStr := "socket out str"
	plugInErr := fmt.Errorf("error")
	plugInErrChan := make(chan error)

	s := &socket{read: socketOutStr}

	p := gcpipe.Plug(s)
	p.SetErrors(plugInErrChan)
	p.Start()

	<-p.Output()

	plugInErrChan <- plugInErr

	p.Close()

	if s.err != plugInErr {
		t.Errorf("expected '%v', got: '%v'", plugInErr, s.err)
	}
}
