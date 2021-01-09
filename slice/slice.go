package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	initLen = 0
	initCap = 3
)

func main() {
	sortedSlice := make([]int, initLen, initCap)
	//var count int = 0
	for true {
		fmt.Print("Enter an integer (X to quit): ")
		var tempStr string
		fmt.Scan(&tempStr)
		if strings.HasPrefix(tempStr, "X") {
			break
		}
		tempInt, err := strconv.Atoi(tempStr)
		if err != nil {
			fmt.Print("That's not an integer; ")
			continue
		}
		//count++
		sortedSlice = append(sortedSlice, tempInt)
		//fmt.Print("Count:", count, "Len:", len(sortedSlice))
		//fmt.Println(sortedSlice)
		for i, val := range sortedSlice {
			if tempInt > val {
				continue
			}
			for j := len(sortedSlice) - 1; j > i; j-- {
				sortedSlice[j] = sortedSlice[j-1]
			}
			sortedSlice[i] = tempInt
			break
		}
		fmt.Println(sortedSlice)

	}
}
