package gccommand_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func TestNewInsert(t *testing.T) {
	c, ok := gccommand.NewInsert("key", "15")

	if !ok {
		t.Error("Insert with integer param should be parsed properly")
	}

	if c.Key() != "key" {
		t.Error("Key should be set properly")
	}

	if c.Item().Value() != "15" {
		t.Error("Value should be parsed properly")
	}
}

func TestInsertExec(t *testing.T) {
	m := gcmap.New()
	c, _ := gccommand.NewInsert("key", "15")

	log, err := c.Exec(m)

	if err != nil {
		t.Error("Insert should work correctly")
	}

	if log != "key := (integer) 15" {
		t.Error("Insert should return correct log")
	}

	i, ok := m.Get("key")

	if !ok {
		t.Error("Insert should insert new value into map")
	}

	if i.Value() != "15" {
		t.Error("Insert should insert correct new value into map")
	}
}
