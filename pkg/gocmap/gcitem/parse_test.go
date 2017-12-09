package gcitem_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
)

func TestParseInteger(t *testing.T) {
	i, ok := gcitem.Parse("10")

	if !ok {
		t.Error("Integer value is expected to be parsed successfully")
	}

	if i.Value() != "10" {
		t.Error("Integer value is expected to be parsed correctly")
	}
}

func TestParseUndefined(t *testing.T) {
	_, ok := gcitem.Parse("aaa")

	if ok {
		t.Error("Undefined value should not be parsed successfully")
	}
}
