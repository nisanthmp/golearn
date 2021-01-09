package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	slice := make([]int, 0, 10)

	fmt.Println("Enter upto 10 integers, separated by spaces, that need to be sorted:")
	br := bufio.NewReader(os.Stdin)

	line, _, err := br.ReadLine()
	if err != nil {
		//fmt.Println("Error reading input")
		panic("Error reading input. quitting...")
	}
	tempSlice := strings.Split(string(line), " ")

	for _, snum := range tempSlice {
		num, err := strconv.Atoi(snum)
		if err != nil {
			panic("A non-integer entered. quitting...")
		}
		slice = append(slice, num)
	}
	bubbleSort(slice)

	for _, num := range slice {
		fmt.Print(num, " ")
	}
	fmt.Print("\n")
}

func bubbleSort(slice []int) {
	length := len(slice)

	for i := 1; i < length; i++ {
		for j := 0; j < length-i; j++ {
			if slice[j] > slice[j+1] {
				swap(slice, j)
			}
		}
	}
}

func swap(slice []int, j int) {
	temp := slice[j]
	slice[j] = slice[j+1]
	slice[j+1] = temp
}
