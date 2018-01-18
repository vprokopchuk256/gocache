package gcpipe

import (
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcsocket"
)

func Pipe(m *gcmap.Map, s gcsocket.Socket) {
	//plug
	p := Plug(s)

	//parser
	cmds, parserErrors := Parser(p.Done(), p.Output())

	//executor
	logs, executorErrors := Executor(m, p.Done(), cmds)
	p.SetInput(logs)

	//errors multiplexer
	errors := Multiplexer(p.Done(), parserErrors, executorErrors)
	p.SetErrors(errors)

	//start
	p.Start()
}
