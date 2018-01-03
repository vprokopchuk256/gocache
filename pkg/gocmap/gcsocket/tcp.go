package gcsocket

import (
	"fmt"
	"net"
)

type tcp struct {
	buff []byte
	conn net.Conn
}

func TCP(conn net.Conn) *tcp {
	return &tcp{
		conn: conn,
		buff: make([]byte, 1024),
	}
}

func (t *tcp) Close() {
	t.conn.Close()
}

func (t *tcp) Read() (string, error) {
	n, err := t.conn.Read(t.buff)

	if err != nil {
		return "", fmt.Errorf("could not read data from the socket: %v", err)
	}

	return string(t.buff[:n]), nil
}

func (t *tcp) Write(s string) error {
	_, err := t.conn.Write([]byte(s))

	if err != nil {
		return fmt.Errorf("could not write data to the socket: %v", err)
	}

	return nil
}

func (t *tcp) Error(e error) error {
	return t.Write(fmt.Sprintf("error: %v", e))
}
