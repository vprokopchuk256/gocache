package gcsource

import "net"

type TCP struct {
	conn   net.Conn
	buff   []byte
	output chan string
	done   chan interface{}
}

func (t *TCP) Output() <-chan string {
	return t.output
}

func (t *TCP) Done() <-chan interface{} {
	return t.done
}

func (t *TCP) Close() {
	t.done <- true
	close(t.done)
	close(t.output)
}

func New(conn net.Conn) *TCP {
	t := &TCP{
		conn:   conn,
		buff:   make([]byte, 1024),
		output: make(chan string),
		done:   make(chan interface{}),
	}

	t.start()

	return t
}

func (t *TCP) read() {
	n, err := t.conn.Read(t.buff)
	if err != nil {
		t.Close()
	} else {
		t.output <- string(t.buff[:n])
	}
}

func (t *TCP) start() {
	for {
		select {
		case <-t.done:
			return
		default:
			t.read()
		}
	}
}
