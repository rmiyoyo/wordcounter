package main

import (
	"bytes"
	"testing"
)

func TestCountWords (t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	res := count(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d.\n", exp, res)
	}
}