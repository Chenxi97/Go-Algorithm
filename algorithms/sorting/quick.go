package sorting

import (
	"math/rand"
)

func partition(arr []int, left, right int) int {
	p := rand.Intn(right-left)+left
	arr[left], arr[p] = arr[p], arr[left]
	temp := arr[left]
	for left < right {
		for left < right && arr[right] >= temp {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= temp {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = temp
	return left
}

func quickSort(arr []int, left, right int) {
	if left < right {
		p := partition(arr, left, right)
		quickSort(arr,left, p-1)
		quickSort(arr,p+1,right)
	}
}

//QuickSort - https://en.wikipedia.org/wiki/Quicksort
func QuickSort(arr []int) {
	quickSort(arr,0,len(arr)-1)
}
