package gccommand_test

import (
	"testing"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gccommand"
)

func TestParseValidIns(t *testing.T) {
	_, err := gccommand.Parse("ins key 10")

	if err != nil {
		t.Fatalf("Expected ins command to be parsed without errors, but got", err)
	}
}

func TestParseInvalidIns(t *testing.T) {
	if _, err := gccommand.Parse("ins"); err == nil {
		t.Error("Expected to return an error while parsing ins without params")
	}
}

func TestParseValidInc(t *testing.T) {
	_, err := gccommand.Parse("inc key")

	if err != nil {
		t.Fatalf("Expected inc command to be parsed without errors")
	}
}

func TestParseValidGet(t *testing.T) {
	_, err := gccommand.Parse("get key")

	if err != nil {
		t.Fatalf("Expected get command to be parsed without errors")
	}
}
