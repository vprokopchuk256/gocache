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

func ParseInteger(value string) (*Integer, error) {
	v, err := strconv.Atoi(strings.TrimSpace(value))

	if err != nil {
		return nil, fmt.Errorf("could not parse integer item: %v", err)
	}

	return NewInteger(v), nil
}

func (i *Integer) Inc() *Integer {
	i.value++

	return i
}

func (i *Integer) Value() string {
	return strconv.Itoa(i.value)
}

func (i *Integer) String() string {
	suffix := "non expirable"
	if i.Expirable() {
		suffix = fmt.Sprintf("expire in %v msc", i.ExpireIn())
	}

	return fmt.Sprintf("(integer) %v, %v", i.Value(), suffix)
}
