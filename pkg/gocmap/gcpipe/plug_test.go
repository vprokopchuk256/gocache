package gcpipe_test

import (
	"fmt"
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcpipe"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/sockettest"
)

func TestPlugSuccess(t *testing.T) {
	socketOutStr := "socket out str"
	plugInStr := "plug in str"
	plugInChan := make(chan string)

	s := &sockettest.Socket{ReadData: socketOutStr}

	p := gcpipe.Plug(s)
	p.SetInput(plugInChan)
	p.Start()

	plugOutStr := <-p.Output()

	plugInChan <- plugInStr

	p.Close()

	if plugOutStr != socketOutStr {
		t.Errorf("expected '%v', got: '%v'", socketOutStr, plugOutStr)
	}

	if s.WriteData != plugInStr {
		t.Errorf("expected '%v', got: '%v'", plugInStr, s.WriteData)
	}
}

func TestPlugError(t *testing.T) {
	socketOutStr := "socket out str"
	plugInErr := fmt.Errorf("error")
	plugInErrChan := make(chan error)

	s := &sockettest.Socket{ReadData: socketOutStr}

	p := gcpipe.Plug(s)
	p.SetErrors(plugInErrChan)
	p.Start()

	<-p.Output()

	plugInErrChan <- plugInErr

	p.Close()

	if s.ErrData != plugInErr {
		t.Errorf("expected '%v', got: '%v'", plugInErr, s.ErrData)
	}
}
