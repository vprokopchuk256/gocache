package gccommand

import (
	"fmt"
	"strings"
)

type parser func(string) (Command, error)

var parsers = map[string]parser{
	"ins":  Insert,
	"inc":  Inc,
	"get":  Get,
	"show": Show,
}

func Parse(sCmd string) (Command, error) {
	cName, params, err := extractCName(sCmd)
	if err != nil {
		return nil, err
	}

	return createCommand(cName, params)
}

func extractCName(sCmd string) (string, string, error) {
	parts := strings.SplitAfterN(sCmd, " ", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("could not extract command name from 'sCmd'")
	}

	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), nil
}

func createCommand(cName string, params string) (Command, error) {
	parser, ok := parsers[cName]
	if !ok {
		return nil, fmt.Errorf("could not recognize command '%v'", cName)
	}

	return parser(params)
}
