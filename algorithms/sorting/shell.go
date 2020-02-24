package sorting

// ShellSort - http://en.wikipedia.org/wiki/Shellsort
func ShellSort(arr []int) {
	for d := len(arr) / 2; d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			temp, j := arr[i], i
			for ; j >= d && arr[j-d] > temp; j -= d {
				arr[j] = arr[j-d]
			}
			arr[j] = temp
		}
	}
}
