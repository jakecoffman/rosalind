package main

func CountNucleotides(dna string) (A int, C int, G int, T int) {
	for i := 0; i < len(dna); i++ {
		switch {
		case dna[i] == 'A':
			A++
		case dna[i] == 'C':
			C++
		case dna[i] == 'G':
			G++
		case dna[i] == 'T':
			T++
		}
	}
	return
}

type nucCount struct {
	A int
	C int
	G int
	T int
}

func CountNucleotidesParallel(dna string) (A int, C int, G int, T int) {
	c := make(chan nucCount)
	l := len(dna)
	go func() {
		A, C, G, T := CountNucleotides(dna[0 : l/4])
		c <- nucCount{A, C, G, T}
	}()
	go func() {
		A, C, G, T := CountNucleotides(dna[l/4 : l/2])
		c <- nucCount{A, C, G, T}
	}()
	go func() {
		A, C, G, T := CountNucleotides(dna[l/2 : l*3/4])
		c <- nucCount{A, C, G, T}
	}()
	go func() {
		A, C, G, T := CountNucleotides(dna[l*3/4:])
		c <- nucCount{A, C, G, T}
	}()
	for i := 0; i < 4; i++ {
		msg := <-c
		A += msg.A
		C += msg.C
		G += msg.G
		T += msg.T
	}
	return
}
