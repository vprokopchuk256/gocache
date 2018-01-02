package server

import (
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gclistener"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func Start() {
	m := gcmap.New()

	//tcp
	gclistener.TCP(m, "localhost", "3333")
}
