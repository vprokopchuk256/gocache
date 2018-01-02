package gcsocket

type Socket interface {
	Read() (string, error)
	Write(string) error
	Close()
}
