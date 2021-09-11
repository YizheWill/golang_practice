package main

import (
	"fmt"
)

func main() {
	// str := "hello world"
	// fmt.Println(strings.Index(str, "h"))
	// str1 := str[0:3]
	// fmt.Println(str1)
	num := 1
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("any")
	}
}
