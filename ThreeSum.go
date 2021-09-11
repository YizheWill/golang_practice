package main

// func twoSum(array []int, target int) [][]int {
// 	res := [][]int{}
// 	resHash := map[int]bool{}
// 	for _, num := range array {
// 		if !resHash[num] && resHash[target-num] {
// 			res = append(res, []int{target - num, num})
// 		}
// 		resHash[num] = true
// 	}
// 	return res
// }

// func addItemToArray(array []int, item int) []int {
// 	first := array[0]
// 	second := array[1]
// 	if first <= item && item <= second {
// 		return []int{first, item, second}
// 	} else if first > item {
// 		return []int{item, first, second}
// 	} else if second < item {
// 		return []int{first, second, item}
// 	} else {
// 		return []int{}
// 	}
// }

// func ThreeNumberSum(array []int, target int) [][]int {
// 	sort.Ints(array)
// 	var res [][]int
// 	for idx, num := range array {
// 		newTarget := target - num
// 		interArray := append(array[:idx], array[idx+1:]...)
// 		twoSumRes := twoSum(interArray, newTarget)
// 		if len(twoSumRes) > 0 {
// 			for _, arr := range twoSumRes {
// 				res = append(res, addItemToArray(arr, newTarget))
// 			}
// 		}
// 	}
// 	return res
// }

func ThreeNumberSum(array []int, target int) [][]int {
	res := [][]int{}
	for idx := 0; idx < len(array)-2; idx++ {
		first := array[idx]
		second := idx + 1
		third := len(array) - 1
		for second < third {
			if first+array[second]+array[third] == target {
				res = append(res, []int{first, second, third})
			} else if first+array[second]+array[third] > target {
				third--
			} else {
				second++
			}
		}
	}
	return res
}
