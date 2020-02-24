package utils

import (
	"math/rand"
	"time"
)

// RandArray returns a n-size integer array
func RandArray(n int) []int {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
