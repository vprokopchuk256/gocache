package gctest

import "fmt"

type socket struct {
	read string
	out  chan string
	errs chan error
	open bool
}

func SocketData(read string) *socket {
	return &socket{
		read: read,
		out:  make(chan string),
		errs: make(chan error),
		open: true,
	}
}

func (s *socket) Read() (string, error) {
	if !s.open {
		return "", fmt.Errorf("could not read from the closed socket")
	}

	return s.read, nil
}

func (s *socket) Write(w string) error {
	if !s.open {
		return fmt.Errorf("could not write to the closed socket")
	}

	s.out <- w

	return nil
}

func (s *socket) Error(err error) error {
	if !s.open {
		return fmt.Errorf("could not write to the closed socket")
	}

	s.errs <- err

	return nil
}

func (s *socket) Close() {
	s.open = false
}

func (s *socket) WaitForOutput() string {
	return <-s.out
}

func (s *socket) WaitForError() error {
	return <-s.errs
}
