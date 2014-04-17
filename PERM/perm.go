package main

import (
	"fmt"
	"sort"

	"github.com/cznic/mathutil"
)

func Permutations(input []int) [][]int {
	result := [][]int{}
	temp := make([]int, len(input))
	copy(temp, input)
	mathutil.PermutationFirst(sort.IntSlice(temp))
	appendme := make([]int, len(input))
	copy(appendme, temp)
	result = append(result, appendme)

	for {
		done := mathutil.PermutationNext(sort.IntSlice(temp))
		if !done {
			return result
		}
		appendme = make([]int, len(input))
		copy(appendme, temp)
		result = append(result, appendme)
	}
}

func main() {
	ps := Permutations([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(len(ps))
	for _, p := range ps {
		for _, i := range p {
			fmt.Printf("%v ", i)
		}
		fmt.Println()
	}
}
