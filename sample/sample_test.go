package main

import (
	"testing"
)

func TestOk(t *testing.T) {
	t.Log("good")
}

func TestNg(t *testing.T) {
	t.Error("failed")
}
