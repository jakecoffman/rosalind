package main

import "testing"

func Test_Perm(t *testing.T) {
	ps := Permutations([]int{1, 2, 3})
	if len(ps) != 6 {
		t.Error("Got", len(ps))
	}
	// TODO: Test the permutations are permuted correctly
}
