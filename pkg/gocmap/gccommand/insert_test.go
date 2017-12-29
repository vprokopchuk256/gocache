package gccommand_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func TestNewInsert(t *testing.T) {
	c, err := gccommand.NewInsert("key", "15")

	if err != nil {
		t.Fatalf("Insert with integer param should be parsed properly")
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

func TestParseInsertValid(t *testing.T) {
	c, err := gccommand.ParseInsert("key 10")

	if err != nil {
		t.Fatalf("error is not expected")
	}

	ins, ok := c.(*gccommand.Insert)

	if !ok {
		t.Fatalf("insert command is expected")
	}

	if ins.Key() != "key" {
		t.Error("parsed key is expected to eq 'key', got", ins.Key())
	}
}
