package sorting

// SelectionSort - http://en.wikipedia.org/wiki/Selection_sort
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		k := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[k] {
				k = j
			}
		}
		arr[k], arr[i] = arr[i], arr[k]
	}
}
