package gccommand_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func TestNewGet(t *testing.T) {
	c, err := gccommand.NewGet("key")

	if err != nil {
		t.Fatalf("get construction always correct")
	}

	if c.Key() != "key" {
		t.Fatalf("key should be set properly")
	}
}

func TestParseGet(t *testing.T) {
	c, err := gccommand.ParseGet("key")

	if err != nil {
		t.Fatalf("get parsing always correct")
	}

	inc, ok := c.(*gccommand.Get)

	if !ok {
		t.Fatalf("get command is expected")
	}

	if inc.Key() != "key" {
		t.Fatalf("key should be set properly")
	}
}

func TestGetExecWithExistingKey(t *testing.T) {
	m := gcmap.New()
	c, _ := gccommand.NewGet("key")
	i := gcitem.NewInteger(10)

	m.Set("key", i)

	log, ok := c.Exec(m)

	if ok != nil {
		t.Fatalf("integer value should be treated without errors")
	}

	if log != "key := (integer) 10" {
		t.Error("get result should be logged properly")
	}
}
