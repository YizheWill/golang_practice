package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	fmt.Println(strings.Index(str, "h"))
	str1 := str[0:3]
	fmt.Println(str1)
}
