package gcpipe_test

import (
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
	expected := "out string"

	s := &socket{read: expected}

	p := gcpipe.Plug(s)
	p.Start()

	got := <-p.Output()

	if got != expected {
		t.Errorf("Expected '%v', got: '%v'", expected, got)
	}
}

// func TestPlugOutputError(t *testing.T) {
// 	s := &socket{readErr: fmt.Errorf("error")}

// 	p := gcpipe.Plug(s)
// 	p.Start()

// 	<-p.Done()
// }

// func TestPlugInput(t *testing.T) {
// 	expected := "in string"

// 	s := &socket{}

// 	p := gcpipe.Plug(s)
// 	p.Start()

// 	p.Input() <- expected

// 	if s.write != expected {
// 		t.Errorf("Expected '%v', got: '%v'", expected, s.write)
// 	}
// }
