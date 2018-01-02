package gcsocket

import (
	"fmt"
	"net"
	"sync"
)

type tcp struct {
	conn   net.Conn
	output chan string
	input  <-chan string
	errors <-chan error
	done   chan interface{}
	wg     sync.WaitGroup
	buff   []byte
}

func TCP(c net.Conn) *tcp {
	return &tcp{
		conn:   c,
		output: make(chan string),
		done:   make(chan interface{}),
		buff:   make([]byte, 1024),
	}
}

func (t *tcp) SetInput(input <-chan string) {
	t.input = input
}

func (t *tcp) SetErrors(errors <-chan error) {
	t.errors = errors
}

func (t *tcp) Start() {
	t.wg.Add(2)

	go func() {
		defer t.wg.Done()
		defer t.close()

		for {
			sCmd, err := t.read()

			if err != nil {
				return
			}

			select {
			case <-t.done:
				return
			case t.output <- sCmd:
			}
		}
	}()

	go func() {
		defer t.wg.Done()

		for {
			select {
			case <-t.done:
				return
			case sCmd := <-t.input:
				t.write(sCmd)
			case err := <-t.errors:
				t.write(err.Error())
			}
		}
	}()

}

func (t *tcp) Output() <-chan string {
	return t.output
}

func (t *tcp) Done() <-chan interface{} {
	return t.done
}

func (t *tcp) close() {
	close(t.done)

	t.wg.Wait()
	t.conn.Close()
}

func (t *tcp) read() (string, error) {
	n, err := t.conn.Read(t.buff)
	if err != nil {
		return "", fmt.Errorf("could not read command: %v", err)
	}

	return string(t.buff[:n]), nil
}

func (t *tcp) write(s string) {
	t.conn.Write([]byte(s))
}
