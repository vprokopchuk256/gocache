package gcpipe

import (
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcsocket"
)

func Pipe(m *gcmap.Map, s gcsocket.Socket) {
	//parser
	cmds, parserErrors := Parser(s.Done(), s.Output())

	//executor
	logs, executorErrors := Executor(m, s.Done(), cmds)
	s.SetInput(logs)

	//errors multiplexer
	errors := Multiplexer(s.Done(), parserErrors, executorErrors)
	s.SetErrors(errors)

	//start
	s.Start()
}
