package main

import "testing"

func Test_Complement(t *testing.T) {
	dna := "AAAACCCGGT"
	comp := Complement([]byte(dna))
	if string(comp) != "ACCGGGTTTT" {
		t.Error(string(comp))
	}
	if string(Complement(comp)) != dna {
		t.Error("Complement not idempotent")
	}
}
