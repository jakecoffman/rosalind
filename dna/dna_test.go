package main

import (
	"bytes"
	"log"
	"math/rand"
	"testing"
)

func Test_CountNucleotides(t *testing.T) {
	dna := "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"
	a, c, g, th := CountNucleotides(dna)
	if a != 20 || c != 12 || g != 17 || th != 21 {
		t.Errorf("Got %v %v %v %v", a, c, g, th)
	}
}

func Test_CountNucleotidesParallel(t *testing.T) {
	dna := "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"
	a, c, g, th := CountNucleotidesParallel(dna)
	if a != 20 || c != 12 || g != 17 || th != 21 {
		t.Errorf("Got %v %v %v %v", a, c, g, th)
	}
}

var dna string

func init() {
	dna = randomDna()
}

func Benchmark_CountNucleotides(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountNucleotides(dna)
	}
}

func Benchmark_CountNucleotidesParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountNucleotidesParallel(dna)
	}
}

func randomDna() string {
	log.Println("Generating random DNA string")
	var buffer bytes.Buffer
	for i := 0; i < 100000000; i++ {
		v := rand.Intn(4)
		switch {
		case v == 0:
			buffer.WriteString("A")
		case v == 1:
			buffer.WriteString("C")
		case v == 2:
			buffer.WriteString("G")
		case v == 3:
			buffer.WriteString("T")
		}
	}
	log.Println("Done generating random DNA string of size", buffer.Len())
	return buffer.String()
}
