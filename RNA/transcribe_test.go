package main

import (
	"bytes"
	"testing"
)

func Test_Transcribe(t *testing.T) {
	dna := "GATGGAACTTGACTACGTAAATT"
	exp := "GAUGGAACUUGACUACGUAAAUU"
	input := bytes.NewBufferString(dna)
	var output bytes.Buffer

	Transcribe(input, &output)

	if output.String() != exp {
		t.Errorf("expected %v (len %v) got %v (len %v)",
			exp, len(exp), output.String(), len(output.String()))
	}
}
