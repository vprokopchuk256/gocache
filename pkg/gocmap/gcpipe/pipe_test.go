package gcpipe_test

// import (
// 	"testing"

// 	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
// 	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcpipe"
// )

// type socket struct {
// 	sCmd   string
// 	output chan string
// 	input  <-chan string
// 	errors <-chan error
// 	done   chan interface{}
// }

// func newSocket(sCmd string) *socket {
// 	return &socket{
// 		sCmd:   sCmd,
// 		output: make(chan string),
// 		done:   make(chan interface{}),
// 	}
// }

// func (s *socket) Output() <-chan string {
// 	return s.output
// }

// func (s *socket) Done() <-chan interface{} {
// 	return s.done
// }

// func (s *socket) SetInput(input <-chan string) {
// 	s.input = input
// }

// func (s *socket) SetErrors(errors <-chan error) {
// 	s.errors = errors
// }

// func (s *socket) close() {
// 	close(s.done)
// }

// func (s *socket) Start() {
// 	go func() {
// 		s.output <- s.sCmd
// 	}()
// }

// func TestPipeSuccess(t *testing.T) {
// 	m := gcmap.New()
// 	s := newSocket("ins key 10")
// 	defer s.close()

// 	gcpipe.Pipe(m, s)

// 	log := <-s.input

// 	if log != "key := (integer) 10" {
// 		t.Errorf("expected log to be correct, got: %v", log)
// 	}
// }

// func TestPipeParsingError(t *testing.T) {
// 	m := gcmap.New()
// 	s := newSocket("some undefined command")
// 	defer s.close()

// 	gcpipe.Pipe(m, s)

// 	err := <-s.errors

// 	if err == nil {
// 		t.Errorf("expected error to be returned")
// 	}
// }

// func TestPipeOperationError(t *testing.T) {
// 	m := gcmap.New()
// 	s := newSocket("inc key")
// 	defer s.close()

// 	gcpipe.Pipe(m, s)

// 	err := <-s.errors

// 	if err == nil {
// 		t.Errorf("expected error to be returned")
// 	}
// }
