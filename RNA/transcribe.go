package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func Transcribe(dna string) string {
	return strings.Replace(dna, "T", "U", -1)
}

func main() {
	bytes, err := ioutil.ReadFile("dataset.dat")
	if err != nil {
		log.Fatalln(err)
	}
	rna := Transcribe(string(bytes))
	ioutil.WriteFile("output.dat", []byte(rna), 0644)
}
