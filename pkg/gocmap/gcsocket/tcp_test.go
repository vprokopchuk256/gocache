package gcsocket_test

import (
	"net"
	"testing"
	"time"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcsocket"
)

type conn struct {
	read  []byte
	write []byte
	open  bool
}

func (c *conn) Read(b []byte) (n int, err error) {
	copy(b, c.read)

	return len(c.read), nil
}

func (c *conn) Write(b []byte) (n int, err error) {
	copy(c.write, b)

	return len(b), nil
}

func (c *conn) Close() error {
	c.open = false

	return nil
}

func (c *conn) LocalAddr() net.Addr {
	return nil
}

func (c *conn) RemoteAddr() net.Addr {
	return nil
}

func (c *conn) SetDeadline(t time.Time) error {
	return nil
}

func (c *conn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *conn) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestRead(t *testing.T) {
	expected := "read string"

	c := &conn{read: []byte("read string")}

	s := gcsocket.TCP(c)
	defer s.Close()

	if sCmd, _ := s.Read(); sCmd != "read string" {
		t.Errorf("expected %v, got", expected)
	}
}

func TestWrite(t *testing.T) {
	expected := "write string"

	c := &conn{write: make([]byte, len(expected))}

	s := gcsocket.TCP(c)
	defer s.Close()

	s.Write(expected)

	if string(c.write) != expected {
		t.Errorf("expected %v, got", expected)
	}
}

func TestClose(t *testing.T) {
	c := &conn{}

	s := gcsocket.TCP(c)

	s.Close()

	if c.open {
		t.Errorf("expected connection to be closed")
	}
}
