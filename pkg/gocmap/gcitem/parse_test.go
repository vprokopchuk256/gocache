package gcitem_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
)

func TestParseInteger(t *testing.T) {
	i, err := gcitem.Parse("10")

	if err != nil {
		t.Error("Integer value is expected to be parsed successfully")
	}

	if i.Value() != "10" {
		t.Error("Integer value is expected to be parsed correctly")
	}
}

func TestParseUndefined(t *testing.T) {
	if _, err := gcitem.Parse("aaa"); err == nil {
		t.Error("Undefined value should not be parsed successfully")
	}
}
