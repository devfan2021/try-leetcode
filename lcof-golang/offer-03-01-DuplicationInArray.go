package main

func findRepeatNumber(nums []int) int {
	vMap := make(map[int]int)
	for _, item := range nums {
		_, exist := vMap[item]
		if exist {
			return item
		}
		vMap[item] = 1
	}
	return -1
}

func findRepeatNumber2(nums []int) int {
	vSet := make(map[int]bool)
	for _, item := range nums {
		value := vSet[item]
		if value {
			return item
		}
		vSet[item] = true
	}
	return -1
}
