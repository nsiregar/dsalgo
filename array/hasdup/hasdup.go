package hasdup

import "errors"

func Verify(arr []int) bool {
	if len(arr) < 2 {
		panic(errors.New("Array length less than 2"))
	}

	for _, val := range arr {
		valCount := Counter(arr, val)
		if valCount > 1 {
			return true
		}
	}
	return false
}

func Counter(arr []int, target int) int {
	var count int
	for idx := 0; idx < len(arr); idx++ {
		if count > 1 {
			return count
		}

		if target == arr[idx] {
			count++
		}
	}
	return count
}
