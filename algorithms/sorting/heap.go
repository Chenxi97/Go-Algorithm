package sorting

func sift(arr []int, i, n int) {
	maxChild := 0
	for i*2+1 < n {
		if i*2+1 == n-1 {
			maxChild = i*2 + 1
		} else if arr[i*2+1] > arr[i*2+2] {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}
		if arr[i] >= arr[maxChild] {
			break
		}
		arr[i], arr[maxChild] = arr[maxChild], arr[i]
		i = maxChild
	}
}

// HeapSort - https://en.wikipedia.org/wiki/Heapsort
func HeapSort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		sift(arr, i, len(arr))
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		sift(arr, 0, i)
	}
}
