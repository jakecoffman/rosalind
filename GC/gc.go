package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type GCData struct {
	id         string
	data       string
	percentage float64
}

type ByPercentage []GCData

func (a ByPercentage) Len() int           { return len(a) }
func (a ByPercentage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPercentage) Less(i, j int) bool { return a[i].percentage > a[j].percentage }

func GC(dna string) float64 {
	gc := 0
	for _, a := range dna {
		if a == 'G' || a == 'C' {
			gc++
		}
	}
	return float64(gc) / float64(len(dna)) * 100.0
}

func LargestGc(filename string) (GCData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return GCData{}, err
	}

	var gc *GCData
	gcs := []GCData{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), ">") {
			if gc != nil {
				gc.percentage = GC(gc.data)
				gcs = append(gcs, *gc)
			}
			gc = &GCData{id: strings.Replace(scanner.Text(), ">", "", 1), data: ""}
		} else {
			gc.data = gc.data + scanner.Text()
		}
	}
	gc.percentage = GC(gc.data)
	gcs = append(gcs, *gc)

	sort.Sort(ByPercentage(gcs))
	return gcs[0], nil
}

func main() {
	p, err := LargestGc("dataset.dat")
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("output.dat", []byte(fmt.Sprintf("%v\n%v", p.id, p.percentage)), 0644)
}
