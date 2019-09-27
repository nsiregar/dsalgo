package twosum

import "errors"

func TwoSum(arr []int, target int) [2]int {
	if len(arr) < 2 {
		panic(errors.New("Required minimum 2 array members"))
	}
	var result [2]int
	arrMap := make(map[int]int)

	for index := 0; index < len(arr); index++ {
		current := arr[index]
		diff := target - current

		if _, ok := arrMap[diff]; !ok {
			arrMap[current] = index
		} else {
			result[0] = arr[arrMap[diff]]
			result[1] = arr[index]
			break
		}
	}
	return result
}
