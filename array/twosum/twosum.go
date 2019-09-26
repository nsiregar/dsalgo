package twosum

func TwoSum(arr []int, target int) [2]int {
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
