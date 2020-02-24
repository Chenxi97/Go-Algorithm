package sorting

// InsertionSort - https://en.wikipedia.org/wiki/Insertion_sort
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp, j := arr[i], i
		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
}
