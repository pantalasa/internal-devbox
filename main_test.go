package main

import "testing"

func TestVersionConst(t *testing.T) {
	if version == "" {
		t.Fatal("version must not be empty")
	}
}
