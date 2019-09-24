package product

func Product(arr []int) []int {
	var productArr []int
	arrLength := len(arr)
	if arrLength == 1 {
		return productArr
	}

	temp := 1
	for idx := 0; idx < arrLength; idx++ {
		productArr = append(productArr, temp)
		temp *= arr[idx]
	}

	temp = 1
	for idx := arrLength - 1; idx >= 0; idx-- {
		productArr[idx] *= temp
		temp *= arr[idx]
	}
	return productArr
}
