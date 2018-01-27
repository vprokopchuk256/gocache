package gcpipe_test

import (
	"strings"
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcpipe"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gctest"
)

func TestPipeSuccess(t *testing.T) {
	m := gcmap.New()
	s := gctest.SocketData("ins key 10")

	gcpipe.Pipe(m, s)

	log := s.WaitForOutput()

	if !strings.Contains(log, "key := (integer) 10") {
		t.Errorf("expected log to be correct, got: %v", log)
	}
}

func TestPipeParsingError(t *testing.T) {
	m := gcmap.New()
	s := gctest.SocketData("some unknown command")

	gcpipe.Pipe(m, s)

	err := s.WaitForError()

	if err == nil {
		t.Errorf("expected error to be returned")
	}
}

func TestPipeOperationError(t *testing.T) {
	m := gcmap.New()
	s := gctest.SocketData("inc key")

	gcpipe.Pipe(m, s)

	err := s.WaitForError()

	if err == nil {
		t.Errorf("expected error to be returned")
	}
}
