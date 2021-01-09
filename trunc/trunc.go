package main

import "fmt"

func main() {
	var inp float32 = 0.0
	var truncInt int
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&inp)
	truncInt = int(inp)
	fmt.Println(truncInt)
}
