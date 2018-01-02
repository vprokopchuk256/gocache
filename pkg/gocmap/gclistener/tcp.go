package gclistener

import (
	"fmt"
	"net"
	"os"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcpipe"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcsocket"
)

func TCP(m *gcmap.Map, host string, port string) {
	l, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}
	defer l.Close()

	fmt.Println("Listening on " + host + ":" + port)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		gcpipe.Pipe(m, gcsocket.TCP(conn))
	}
}
