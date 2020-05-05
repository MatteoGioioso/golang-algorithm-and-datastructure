package main

import "fmt"

func sort(set []int, key int) []int {
	if key == len(set) {
		return set
	}

	currentSet := set[:key+1]
	for index, _ := range currentSet {
		setLength := len(currentSet)
		keyNum := currentSet[setLength-1-index]
		prevNum := currentSet[setLength-index-2]

		if prevNum > keyNum {
			//Swap the two
			currentSet[setLength-1-index] = prevNum
			currentSet[setLength-index-2] = keyNum
		}

		if setLength-index-2 == 0 {
			break
		}
	}

	return sort(set, key+1)
}

func InsertionSort() {
	unsortedSet := []int{67, 5, 3, 8, 19, 1, 25, 2, 34}
	sortedSet := sort(unsortedSet, 1)
	fmt.Println(sortedSet)
}
