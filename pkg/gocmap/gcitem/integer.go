package gcitem

import (
	"fmt"
	"strconv"
	"strings"
)

type Integer struct {
	Base
	value int
}

func NewInteger(value int) *Integer {
	return &Integer{base(), value}
}

func ParseInteger(value string) (*Integer, bool) {
	v, err := strconv.Atoi(strings.TrimSpace(value))

	if err != nil {
		return nil, false
	}

	return NewInteger(v), true
}

func (i *Integer) Inc() *Integer {
	i.value++

	return i
}

func (i *Integer) Value() string {
	return strconv.Itoa(i.value)
}

func (i *Integer) String() string {
	return fmt.Sprintf("(integer) %v", i.Value())
}
