package main

import (
	"bytes"
	"testing"
)

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("Line1\nLine2\nLine3\n")
	exp := 3
	res := count(b, true, false, false)

	if res != exp {
		t.Errorf("Expected %d lines, got %d.\n", exp, res)
	}
}

func TestCountWords (t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	res := count(b, false, true, false)

	if res != exp {
		t.Errorf("Expected %d, got %d.\n", exp, res)
	}
}

func TestCountChars (t *testing.T) {
	b := bytes.NewBufferString("Test Chars\n")
	exp := 10
	res := count(b, false, false, true)

	if res != exp {
		t.Errorf("Expected %d, got %d.\n", exp, res)
	}
}