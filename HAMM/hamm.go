package main

import "log"

func HAMM(dna1 string, dna2 string) int {
	if len(dna1) != len(dna2) {
		log.Fatal("DNA lengths are not equal")
	}
	hamm := 0
	for i := 0; i < len(dna1); i++ {
		if dna1[i] != dna2[i] {
			hamm++
		}
	}
	return hamm
}
