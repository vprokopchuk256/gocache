package gccommand_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func TestInsertFInteger(t *testing.T) {
	m := gcmap.New()

	c, err := gccommand.Insert("key 15")

	if err != nil {
		t.Fatalf("insert command is expected")
	}

	log, err := c(m)

	if err != nil {
		t.Fatalf("insert should work correctly")
	}

	if log != "key := (integer) 15" {
		t.Error("insert should return correct log")
	}

	i, ok := m.Get("key")

	if !ok {
		t.Error("insert should insert new value into map")
	}

	if i.Value() != "15" {
		t.Error("insert should insert correct new value into map")
	}
}
