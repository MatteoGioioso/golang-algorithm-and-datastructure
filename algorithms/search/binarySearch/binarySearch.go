package binarySearch

/**
Average performance	O(log n)
*/
func search(sortedSet []int, numberToSearch int) bool {
	midPointIndex := len(sortedSet) / 2
	midPointNumber := sortedSet[midPointIndex]

	if midPointNumber == numberToSearch {
		return true
	}

	if len(sortedSet) == 1 {
		return false
	}

	if midPointNumber < numberToSearch {
		return search(sortedSet[midPointIndex:], numberToSearch)
	} else {
		return search(sortedSet[0:midPointIndex], numberToSearch)
	}
}

func BinarySearch(numberToSearch int) bool {
	sortedSet := []int{1, 2, 3, 5, 8, 19, 25, 34, 67}
	return search(sortedSet, numberToSearch)
}
