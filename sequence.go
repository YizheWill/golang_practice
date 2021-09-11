// package main

// import "fmt"

// func IsValidSubsequence(array []int, sequence []int) bool {
// 	// Write your code here.
// 	j := 0
// 	for i := 0; i < len(array) && j < len(sequence); i++ {
// 		if array[i] == sequence[j] {
// 			j++
// 		}
// 		if j == len(sequence) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	array := []int{1, 2, 3, 4, 5, 7}
// 	sequence := []int{2, 3}
// 	res := IsValidSubsequence(array, sequence)
// 	fmt.Println(res)
// }
package main

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	for i, j := 0, 0; i < len(array) && j < len(sequence); i++ {
		if array[i] == sequence[j] {
			j++
			if j == len(sequence) {
				return true
			}
		}
	}
	return false
}
