package searching

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11}
	args := [][]int{{0, -1}, {12, -1}, {5, 2}, {6, -1}, {1, 0}, {11, 5}}
	for i := 0; i < len(args); i++ {
		if res := BinarySearch(arr, args[i][0]); res != args[i][1] {
			t.Errorf("find %v ,want %v got %v", args[i][0], args[i][1], res)
		}
	}
}
func TestLowerBound(t *testing.T){
	arr := []int{1, 3, 5, 7, 9, 11}
	args := [][]int{{0, 0}, {1, 0}, {5, 2}, {6, 3}, {11, 5}, {13, 6}}
	for i := 0; i < len(args); i++ {
		if res := LowerBound(arr, args[i][0]); res != args[i][1] {
			t.Errorf("find %v ,want %v got %v", args[i][0], args[i][1], res)
		}
	}
}

func TestUpperBound(t *testing.T){
	arr := []int{1, 3, 5, 7, 9, 11}
	args := [][]int{{0, 0}, {1, 1}, {5, 3}, {6, 3}, {10, 5}, {13, 6}}
	for i := 0; i < len(args); i++ {
		if res := UpperBound(arr, args[i][0]); res != args[i][1] {
			t.Errorf("find %v ,want %v got %v", args[i][0], args[i][1], res)
		}
	}
}
