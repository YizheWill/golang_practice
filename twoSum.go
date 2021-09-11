package main

import "fmt"

func twoSum1(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == target {
				return []int{array[i], array[j]}
			}
		}
	}
	return []int{}
}

func twoSum2(array []int, target int) []int {
	dict := map[int]bool{}
	for _, val := range array {
		possibleAnswer := target - val
		if dict[possibleAnswer] {
			return []int{possibleAnswer, val}
		}
		dict[val] = true
	}
	return []int{}
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	for i, num := range array {
		fmt.Println("this is i", i)
		fmt.Println("this is num", num)
	}
	for a, b := range "happy" {
		fmt.Println(a, b)
	}
	// 注意：map并没有顺序，iterate的时候并不是按照创建顺序逐个跑的。
	strMap := map[string]string{"a": "a", "b": "b", "c": "c", "d": "d"}
	for k, v := range strMap {
		fmt.Println(k, v)
	}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 22}
	fmt.Println(twoSum1(arr, 23))
	fmt.Println(twoSum2(arr, 23))
}
