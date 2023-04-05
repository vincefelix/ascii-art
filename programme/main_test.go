package main

import (
	"testing"
)

func TestIsPrintable(t *testing.T) {
	if IsPrintable("yuuuy") != true {
		t.Error("expected true")
	} else {
		t.Log("Succeedeed")
	}
}

func TestNewline(t *testing.T) {
	if Newline([]string{"erth", "er", "ert"}) == nil {
		t.Error("expected something")
	} else {
		t.Log("Succeedeed")
	}
}

func TestPrintres(t *testing.T) {
	var slice = [][]string{
		{"a", "a", "a", "a", "a", "a", "a", "a"}, {"1", "2", "3", "4", "5", "6", "7", "8"}}
	if printres(slice) != "a1\na2\na3\na4\na5\na6\na7\na8" {
		t.Error("expected the right result")
	} else {
		t.Log("Succeedeed")
	}

}
