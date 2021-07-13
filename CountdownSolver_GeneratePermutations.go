package main

import (
	"strconv"
)

var permutations = make(map[string][]int64)

func generatePermutations(inputs []int64) [][]int64 {
	heapPermutations(inputs, len(inputs))
	return turnPermutationsIntoSlice(permutations)
}

func turnPermutationsIntoSlice(permutations map[string][]int64) [][]int64 {
	var output [][]int64
	for _, element := range permutations {
		output = append(output, element)
	}
	return output
}

func heapPermutations(inputs []int64, size int) {
	if size == 1 {
		mapString := ""
		var outArr []int64
		for index := range inputs {
			mapString += strconv.FormatInt(inputs[index], 10) + "_"
			outArr = append(outArr, inputs[index])
		}
		permutations[mapString] = outArr
		return
	}
	for i := 0; i < size; i++ {
		heapPermutations(inputs, size-1)
		if size%2 == 1 {
			inputs[0], inputs[size-1] = inputs[size-1], inputs[0]
		} else {
			inputs[i], inputs[size-1] = inputs[size-1], inputs[i]
		}
	}
}
