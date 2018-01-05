package gcpipe

import (
	"sync"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcsocket"
)

type plug struct {
	socket gcsocket.Socket
	output chan string
	input  <-chan string
	errors <-chan error
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

func (p *plug) SetInput(input <-chan string) {
	p.input = input
}

func (p *plug) SetErrors(errors <-chan error) {
	p.errors = errors
}

func (p *plug) Done() <-chan interface{} {
	return p.done
}

func (p *plug) Start() {
	p.wg.Add(1)

	go func() {
		defer p.wg.Done()

		for {
			sCmd, err := p.socket.Read()

			if err != nil {
				go p.Close()
				break
			}

		Out:
			for {
				select {
				case p.output <- sCmd:
					break Out
				case <-p.done:
					return
				}
			}

		In:
			for {
				select {
				case in := <-p.input:
					p.socket.Write(in)
					break In
				case err := <-p.errors:
					p.socket.Error(err)
					break In
				case <-p.done:
					return
				}
			}
		}
	}()
}

func (p *plug) Close() {
	close(p.done)

	p.wg.Wait()

	close(p.output)
	p.socket.Close()
}
