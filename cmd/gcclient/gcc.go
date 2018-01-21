package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/abiosoft/ishell"
)

var conn, _ = net.Dial("tcp", "127.0.0.1:3333")

func main() {
	shell := ishell.New()

	shell.Println("GoCache interactive client")

	cmds(shell, [][]string{
		[]string{"ins", "insert value into the cache"},
		[]string{"inc", "increase value with the specified key"},
	})

	shell.Run()
}

func cmds(shell *ishell.Shell, sCmds [][]string) {
	for _, cmdData := range sCmds {
		cmd(shell, cmdData[0], cmdData[1])
	}
}

func cmd(shell *ishell.Shell, sCmd string, help string) {
	shell.AddCmd(&ishell.Cmd{
		Name: sCmd,
		Help: help,
		Func: runCmd(sCmd),
	})
}

func runCmd(sCmd string) func(*ishell.Context) {
	return func(c *ishell.Context) {
		fmt.Fprintf(conn, fmt.Sprintf("%v %v", sCmd, strings.Join(c.Args, " ")))

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)
	}
}
