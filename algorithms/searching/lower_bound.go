package searching

//LowerBound returns position of the first number >= target
func LowerBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + ((right - left) >> 1)
		if arr[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
