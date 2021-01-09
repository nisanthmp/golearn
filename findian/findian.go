package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Enter a string: ")
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println("That's not a string")
		return
	}
	str = strings.ToLower(str)
	b1 := strings.HasPrefix(str, "i")
	b2 := strings.HasSuffix(str, "n")
	b3 := strings.Contains(str, "a")

	if b1 && b2 && b3 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
