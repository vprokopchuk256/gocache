package gcsocket_test

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcsocket"
)

type conn struct {
	read  []byte
	write []byte
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

func newConn() *conn {
	return &conn{
		read:  []byte("read string"),
		write: make([]byte, 12),
	}
}

func TestTCPOutput(t *testing.T) {
	conn := newConn()

	tcp := gcsocket.TCP(conn)

	tcp.Start()

	out := <-tcp.Output()

	if out != "read string" {
		t.Errorf("expected to put read string into output channel, got: %v", out)
	}
}

func TestTCPInput(t *testing.T) {
	conn := newConn()

	tcp := gcsocket.TCP(conn)

	in := make(chan string)
	defer close(in)

	tcp.SetInput(in)
	tcp.Start()

	in <- "write string"

	if string(conn.write) != "write string" {
		t.Errorf("expected to write output string into channel, got written: %v", string(conn.write))
	}
}

func TestTCPError(t *testing.T) {
	conn := newConn()

	tcp := gcsocket.TCP(conn)

	err := make(chan error)
	defer close(err)

	tcp.SetErrors(err)
	tcp.Start()

	err <- fmt.Errorf("error string")

	if string(conn.write) != "error string" {
		t.Errorf("expected to write error string into channel, got written: %v", string(conn.write))
	}
}
