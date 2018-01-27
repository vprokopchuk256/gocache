package gccommand_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func TestIncFExecWithExistingIntegerKey(t *testing.T) {
	m := gcmap.New()
	i := gcitem.NewInteger(10)

	m.Set("key", i)

	c, err := gccommand.Inc("key")

	if err != nil {
		t.Fatalf("inc parsing always correct")
	}

	log, ok := c(m)

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
