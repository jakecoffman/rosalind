package main

import "strings"

func CountNucleotides(dna string) (A int, C int, G int, T int) {
	A = strings.Count(dna, "A")
	C = strings.Count(dna, "C")
	G = strings.Count(dna, "G")
	T = strings.Count(dna, "T")
	return
}

type nucCount struct {
	Nuc   string
	Count int
}

func CountNucleotidesParallel(dna string) (A int, C int, G int, T int) {
	c := make(chan nucCount)
	go func() {
		c <- nucCount{"A", strings.Count(dna, "A")}
	}()
	go func() {
		c <- nucCount{"C", strings.Count(dna, "C")}
	}()
	go func() {
		c <- nucCount{"G", strings.Count(dna, "G")}
	}()
	go func() {
		c <- nucCount{"T", strings.Count(dna, "T")}
	}()
	for i := 0; i < 4; i++ {
		msg := <-c
		switch {
		case msg.Nuc == "A":
			A = msg.Count
		case msg.Nuc == "C":
			C = msg.Count
		case msg.Nuc == "G":
			G = msg.Count
		case msg.Nuc == "T":
			T = msg.Count
		}

	}
	return
}
