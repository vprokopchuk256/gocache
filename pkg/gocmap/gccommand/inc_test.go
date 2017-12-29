package gccommand_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func TestNewInc(t *testing.T) {
	c, err := gccommand.NewInc("key")

	if err != nil {
		t.Fatalf("inc construction always correct")
	}

	if c.Key() != "key" {
		t.Fatalf("key should be set properly")
	}
}

func TestParseInc(t *testing.T) {
	c, err := gccommand.ParseInc("key")

	if err != nil {
		t.Fatalf("inc parsing always correct")
	}

	inc, ok := c.(*gccommand.Inc)

	if !ok {
		t.Fatalf("inc command is expected")
	}

	if inc.Key() != "key" {
		t.Fatalf("key should be set properly")
	}
}

func TestIncExecWithExistingIntegerKey(t *testing.T) {
	m := gcmap.New()
	c, _ := gccommand.NewInc("key")
	i := gcitem.NewInteger(10)

	m.Set("key", i)

	log, ok := c.Exec(m)

	if ok != nil {
		t.Fatalf("integer value should be treated without errors")
	}

	if log != "key := (integer) 11" {
		t.Error("inc result should be logged properly")
	}

	if i.Value() != "11" {
		t.Error("integer value must be increased")
	}
}
