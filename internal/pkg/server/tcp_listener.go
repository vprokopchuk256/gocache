package server

import (
	"fmt"
	"net"
	"os"
)

func StartTCP(host string, port string) {
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

		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	buff := make([]byte, 1024)

	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	conn.Write([]byte("Message received"))
	conn.Close()
}
