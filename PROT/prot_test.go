package main

import (
	"bytes"
	"testing"
)

func TestProt(t *testing.T) {
	input := "AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA"
	expected := "MAMAPRTEINSTRING"
	inBuf := bytes.NewBufferString(input)
	outBuf := bytes.Buffer{}
	prot(inBuf, &outBuf)
	if outBuf.String() != expected {
		t.Errorf("expected %v got %v", expected, outBuf.String())
	}
}
