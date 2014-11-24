package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
)

func init() {
	rawCodon, err := os.Open("rna-codon.json")
	if err != nil {
		log.Fatal(err)
	}
	translator = map[string]string{}
	err = json.NewDecoder(rawCodon).Decode(&translator)
	if err != nil {
		log.Fatal(err)
	}
}

var translator map[string]string

func main() {
	inFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	if err = prot(inFile, outFile); err != nil {
		log.Fatal(err)
	}
}

// custom scanner that transforms RNA to Protien on the fly
func ScanRNAtoProt(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) < 3 {
		return 0, nil, nil
	}
	rna := string(data[0:3])
	prot, ok := translator[rna]
	if !ok {
		return 0, nil, errors.New("RNA input invalid")
	}
	return 3, []byte(prot), nil
}

func prot(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanRNAtoProt)
	for scanner.Scan() {
		w.Write(scanner.Bytes())
	}
	return scanner.Err()
}
