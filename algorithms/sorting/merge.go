package sorting

import utils "github.com/Chenxi97/Go-Algorithm"

func merge(arr []int, l1, r1, l2, r2 int) {
	temp := make([]int, len(arr))
	i, j, idx := l1, l2, 0
	for i <= r1 && j <= r2 {
		if arr[i] < arr[j] {
			temp[idx] = arr[i]
			idx, i = idx+1, i+1
		} else {
			temp[idx] = arr[j]
			idx, j = idx+1, j+1
		}
	}
	for i <= r1 {
		temp[idx] = arr[i]
		idx, i = idx+1, i+1
	}
	for j <= r2 {
		temp[idx] = arr[j]
		idx, j = idx+1, j+1
	}
	for i = l1; i <= r2; i++ {
		arr[i] = temp[i]
	}
}

// MergeSort http://en.wikipedia.org/wiki/Merge_sort
func MergeSort(arr []int) {
	n := len(arr)
	for step := 2; step/2 < n; step *= 2 {
		for i := 0; i < n; i += step {
			mid := i + step/2 - 1
			if mid+1 < n {
				merge(arr, i, mid, mid+1, utils.IntMin(i+step-1, n-1))
			}
		}
	}
}
