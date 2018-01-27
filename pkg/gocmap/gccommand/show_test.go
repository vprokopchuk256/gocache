package gccommand_test

import (
	"strings"
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func TestShowFExecWithExistingKey(t *testing.T) {
	m := gcmap.New()
	i := gcitem.NewInteger(10)

	m.Set("key", i)

	c, ok := gccommand.Show("key")

	if ok != nil {
		t.Fatalf("get parsing always correct")
	}

	log, ok := c(m)

	if ok != nil {
		t.Fatalf("integer value should be treated without errors")
	}

	if !strings.Contains(log, "(integer) 10") {
		t.Error("insert should return correct log")
	}
}
