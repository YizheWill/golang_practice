package main

import "fmt"

// import (
// 	"fmt"
// )

// func main() {
// 	// str := "hello world"
// 	// fmt.Println(strings.Index(str, "h"))
// 	// str1 := str[0:3]
// 	// fmt.Println(str1)
// 	num := 1
// 	switch num {
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	default:
// 		fmt.Println("any")
// 	}
// }

// import "strings"
// func reversePrefix(word string, ch byte) string {
// 	idx := strings.Index(word, string(ch))
//     if idx <= 0 || len(word) <= 1{
// 		return word
// 	}

// 	res := make([]rune, len(word))
// 	res = []rune(word)
//     for i, j := 0, idx; i < (idx + 1) / 2; i, j = i+1, j-1 {
//         res[i], res[idx] = res[idx], res[i]
//     }

// 	return string(res)
// }
func interchangeableRectangles(rectangles [][]int) int64 {
	m := map[float64]int{}
	for _, item := range rectangles {
		arrKey := float64(item[0]) / float64(item[1])
		m[arrKey] += 1
	}
	sum := 0
	for _, v := range m {
		if v > 1 {
			sum += sumUp(v)
		}
	}
	return int64(sum)
}

func sumUp(num int) int {
	res := 0
	num--
	for num > 0 {
		res += num
		num--
	}
	return res
}
func main() {
	a := [][]int{{4, 8}, {3, 6}, {10, 20}, {15, 30}}
	fmt.Println(interchangeableRectangles(a))
}
