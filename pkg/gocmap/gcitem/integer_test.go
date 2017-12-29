package gcitem_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
)

func TestIntegerNew(t *testing.T) {
	if i := gcitem.NewInteger(10); i.Value() != "10" {
		t.Error("Initially integer value is expected to be set in construction")
	}
}

func TestIntegerInc(t *testing.T) {
	i := gcitem.NewInteger(10)

	v := i.Inc()

	if v.Value() != "11" {
		t.Error("Is expected to increase integer value by 1")
	}
}

func TestIntegerShow(t *testing.T) {
	if i := gcitem.NewInteger(10); i.String() != "(integer) 10" {
		t.Error("Is expected to has correct display value")
	}
}

func TestIntegerParse(t *testing.T) {
	i, err := gcitem.ParseInteger("10")

	if err != nil {
		t.Error("It should successfully parse corrent integer value")
	}

	if i.Value() != "10" {
		t.Error("Correct integer value is expected to be set in construction")
	}
}

func TestIntegerParseSpaces(t *testing.T) {
	i, err := gcitem.ParseInteger(" 10 ")

	if err != nil {
		t.Error("It should successfully parse corrent integer value and ignore leading and trailing spaces")
	}

	if i.Value() != "10" {
		t.Error("Correct integer value is expected to be set in construction")
	}
}

func TestIntegerParseSpacesOnly(t *testing.T) {
	if _, err := gcitem.ParseInteger(" "); err == nil {
		t.Error("It should not parse empty string")
	}
}

func TestIntegerParseInvalidSpaces(t *testing.T) {
	if _, err := gcitem.ParseInteger("aa"); err == nil {
		t.Error("It should not parse non digits")
	}
}
