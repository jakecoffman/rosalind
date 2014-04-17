package main

import (
	"sort"
	"testing"
)

func Test_GC(t *testing.T) {
	d1 := "CCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCCTCCCACTAATAATTCTGAGG"
	p1 := GC(d1)
	d2 := "CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCTATATCCATTTGTCAGCAGACACGC"
	p2 := GC(d2)
	d3 := "CCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGACTGGGAACCTGCGGGCAGTAGGTGGAAT"
	p3 := GC(d3)
	percentages := []float64{p1, p2, p3}
	sort.Float64s(percentages)
	if !aboutEqual(percentages[2], 60.919540, 0.001) {
		t.Errorf("%#v", percentages)
	}
}

func Test_GC2(t *testing.T) {
	p := GC("AGCTATAG")
	if !aboutEqual(p, 37.5, 0.001) {
		t.Errorf("%#v", p)
	}
}

func Test_LargestGc(t *testing.T) {
	v, err := LargestGc("test.dat")
	if err != nil || !aboutEqual(v.percentage, 60.919540, 0.001) {
		t.Errorf("%#v", v)
	}
}

func aboutEqual(val1 float64, val2 float64, thresh float64) bool {
	if val1 >= val2-thresh && val1 <= val2+thresh {
		return true
	}
	return false
}
