package main

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

func Transcribe(input io.Reader, output io.Writer) error {
	buffer := make([]byte, 32*1024)
	for {
		m, err := input.Read(buffer)
		if err != nil {
			return err
		}
		if m == 0 {
			break
		}
		rna := strings.Replace(string(buffer), "T", "U", -1)
		// only write the bytes written or len(output) < len(input)
		n, err := output.Write([]byte(rna[0:m]))
		if err != nil {
			return err
		}
		if n == 0 {
			return errors.New("Failed to write to output in Transcribe")
		}
	}
	return nil
}

func main() {
	file, err := os.Open("dataset.dat")
	if err != nil {
		log.Fatalln(err)
	}
	output, err := os.Create("output.dat")
	if err != nil {
		log.Fatalln(err)
	}
	Transcribe(file, output)
}
