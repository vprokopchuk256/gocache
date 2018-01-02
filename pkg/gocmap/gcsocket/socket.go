package gcsocket

type Socket interface {
	Output() <-chan string
	Done() <-chan interface{}

	SetInput(<-chan string)
	SetErrors(<-chan error)

	Start()
}
