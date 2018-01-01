package gcsource

type Source interface {
	Output() <-chan string
	Done() <-chan interface{}
	Close()
}
