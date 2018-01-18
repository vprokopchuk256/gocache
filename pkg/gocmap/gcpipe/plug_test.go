package gcpipe_test

import (
	"fmt"
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcpipe"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gctest"
)

func TestPlugSuccess(t *testing.T) {
	socketData := "socket out str"
	plugInStr := "plug in str"
	plugInChan := make(chan string)

	s := gctest.SocketData(socketData)

	p := gcpipe.Plug(s)
	p.SetInput(plugInChan)
	p.Start()

	plugOutStr := <-p.Output()
	plugInChan <- plugInStr

	socketInData := s.WaitForOutput()

	if plugOutStr != socketData {
		t.Errorf("expected '%v', got: '%v'", socketData, plugOutStr)
	}

	if socketInData != plugInStr {
		t.Errorf("expected '%v', got: '%v'", plugInStr, socketInData)
	}
}

func TestPlugError(t *testing.T) {
	socketData := "socket out str"
	plugInErr := fmt.Errorf("error")
	plugInErrChan := make(chan error)

	s := gctest.SocketData(socketData)

	p := gcpipe.Plug(s)
	p.SetErrors(plugInErrChan)
	p.Start()

	<-p.Output()

	plugInErrChan <- plugInErr

	socketInError := s.WaitForError()

	if socketInError != plugInErr {
		t.Errorf("expected '%v', got: '%v'", plugInErr, socketInError)
	}
}
