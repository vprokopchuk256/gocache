package gccommand_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"
)

func TestParseValidIns(t *testing.T) {
	c, err := gccommand.Parse("ins key 10")

	if err != nil {
		t.Fatalf("Expected ins command to be parsed without errors, but got", err)
	}

	ins, succ := c.(*gccommand.Insert)

	if !succ {
		t.Error("Expected insert command")
	}

	if ins.Key() != "key" {
		t.Error("Expected parsed key")
	}
}

func TestParseInvalidIns(t *testing.T) {
	if _, err := gccommand.Parse("ins"); err == nil {
		t.Error("Expected to return an error while parsing ins without params")
	}
}

func TestParseValidInc(t *testing.T) {
	c, err := gccommand.Parse("inc key")

	if err != nil {
		t.Fatalf("Expected inc command to be parsed without errors")
	}

	ins, succ := c.(*gccommand.Inc)

	if !succ {
		t.Error("Expected inc command")
	}

	if ins.Key() != "key" {
		t.Error("Expected parsed key")
	}
}
