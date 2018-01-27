package gccommand_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func TestGetFExecWithExistingKey(t *testing.T) {
	m := gcmap.New()
	i := gcitem.NewInteger(10)

	m.Set("key", i)

	c, ok := gccommand.Get("key")

	if ok != nil {
		t.Fatalf("get parsing always correct")
	}

	log, ok := c(m)

	if ok != nil {
		t.Fatalf("integer value should be treated without errors")
	}

	if log != "key := (integer) 10" {
		t.Error("get result should be logged properly")
	}
}
