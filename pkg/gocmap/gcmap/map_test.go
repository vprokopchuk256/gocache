package gcmap_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func TestGetExistingKey(t *testing.T) {
	m := gcmap.New()
	item := gcitem.NewInteger(10)

	m.Set("key", item)

	i, ok := m.Get("key")

	if !ok {
		t.Error("Existing item should be detected")
	}

	if i != item {
		t.Error("Existing item should be returned")
	}
}

func TestGetExistingExpiredKey(t *testing.T) {
	m := gcmap.New()
	item := gcitem.NewInteger(10)
	item.SetExpireIn(0)

	m.Set("key", item)

	_, ok := m.Get("key")

	if ok {
		t.Error("Expired item should be detected")
	}
}

func TestGetNonExistingKey(t *testing.T) {
	m := gcmap.New()

	_, ok := m.Get("key")

	if ok {
		t.Error("Existing item should not be detected")
	}
}

func TestDelete(t *testing.T) {
	m := gcmap.New()

	m.Set("key", gcitem.NewInteger(10))
	m.Delete("key")

	_, ok := m.Get("key")

	if ok {
		t.Error("Deleted item should not be detected")
	}
}
