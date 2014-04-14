package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func Complement(dna []byte) []byte {
	l := len(dna)
	str := make([]string, l)

	for i, c := range dna {
		switch {
		case c == 'A':
			str[l-i-1] = "T"
		case c == 'T':
			str[l-i-1] = "A"
		case c == 'C':
			str[l-i-1] = "G"
		case c == 'G':
			str[l-i-1] = "C"
		}
	}
	return []byte(strings.Join(str, ""))
}

func main() {
	bytes, err := ioutil.ReadFile("dataset.dat")
	if err != nil {
		log.Fatalln(err)
	}
	rna := Complement(bytes)

	ioutil.WriteFile("output.dat", rna, 0644)
}
