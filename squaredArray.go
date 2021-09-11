// // package main
// // func SortedSquaredArray(array []int) []int {
// // 	// Write your code here.
// // 	res := []int{}
// // 	for _, num := range array {
// // 		res = append(res, num*num)
// // 	}
// // 	return res
// // }

// package main
// import "math"
// func squareArray(array []int, reversed bool, square bool) []int {
// 	res := make([]int, len(array))
// 	arrlen := len(array)
// 	for idx, num := range array {
// 		if reversed {
// 			idx = arrlen - 1 - idx
// 		}
// 		if square {
// 			num = num * num
// 		}
// 		res[idx] = num
// 	}
// 	return res
// }

// func SortedSquaredArray(array []int) []int {
// 	// Write your code here.
// 	res := make([]int, len(array))
// 	if len(array) < 2 {
// 		return []int{}
// 	}
// 	var mid int
// 	if array[0] < 0 && array[0]*array[len(array)-1] < 0 {
// 		for idx := 0; idx < len(array); idx++ {
// 			if array[idx] < 0 && idx < len(array)-1 && array[idx+1] >= 0 {
// 				mid = idx
// 			}
// 		}
// 		newArray := squareArray(array[:mid], true, false)
// 		array = array[mid+1 : len(array)]
// 		i, j, k := 0, 0, 0
// 		for i < len(newArray) && j < len(array) {
// 			if (math.Abs(newArray[i]) < math.Abs(array[j])) {

// 			}
// 		}

// 	} else if array[0] < 0 && array[len(array)-1] <= 0 {
// 		return squareArray(array, true, true)
// 	}
// 	return squareArray(array, false, true)
// }

// package main

// func abs(num int) int {
// 	if num < 0 {
// 		return -num
// 	}
// 	return num
// }

// func SortedSquaredArray(array []int) []int {
// 	res := make([]int, len(array))
// 	start := 0
// 	end := len(array) - 1
// 	idx := len(array) - 1
// 	for start < end {
// 		if abs(array[start]) > abs(array[end]) {
// 			res[idx] = array[start] * array[start]
// 			end--
// 		} else {
// 			res[idx] = array[end] * array[end]
// 			start++
// 		}
// 	}

// 	return res
// }

package main

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func SortedSquaredArray(array []int) []int {
	res := make([]int, len(array))
	start := 0
	end := len(array) - 1
	for idx := len(array) - 1; idx >= 0; idx-- {
		if abs(array[start]) < abs(array[end]) {
			res[idx] = array[end] * array[end]
			end--
		} else {
			res[idx] = array[start] * array[start]
			start++
		}
	}
	return res
}
