package main

import "testing"

func Test_Transcribe(t *testing.T) {
	dna := "GATGGAACTTGACTACGTAAATT"
	rna := Transcribe(dna)
	if rna != "GAUGGAACUUGACUACGUAAAUU" {
		t.Error("Failed to transcribe")
	}
}
