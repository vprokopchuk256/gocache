package gcpipe

import (
	"sync"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcsocket"
)

type plug struct {
	socket gcsocket.Socket
	output chan string
	input  <-chan string
	done   chan interface{}
	wg     sync.WaitGroup
}

func Plug(s gcsocket.Socket) *plug {
	return &plug{
		socket: s,
		output: make(chan string),
		done:   make(chan interface{}),
	}
}

func (p *plug) Output() <-chan string {
	return p.output
}

func (p *plug) Input() chan<- string {
	return p.output
}

func (p *plug) Done() <-chan interface{} {
	return p.done
}

func (p *plug) Start() {
	go func() {
		defer p.close()

		for {
			sCmd, err := p.socket.Read()

			if err != nil {
				return
			}

			p.output <- sCmd
		}
	}()
}

func (p *plug) close() {
	close(p.done)
	close(p.output)
	p.socket.Close()
}
