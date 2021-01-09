package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Enter an array of integers, on a line, separated by spaces:")
	br := bufio.NewReader(os.Stdin)
	bline, _, err := br.ReadLine()
	if err != nil {
		panic("Error reading input line")
	}
	sline := string(bline)
	sline = strings.Trim(sline, " \n")
	snums := strings.Split(sline, " ")
	lenNums := len(snums)
	nums := make([]int, 0, lenNums)
	for _, snum := range snums {
		num, err := strconv.Atoi(snum)
		if err != nil {
			fmt.Println("snum =", snum, "err:", err)
			panic("Did not get integers. Quitting...")
		}
		nums = append(nums, num)
	}
	wg := sync.WaitGroup{}
	quarter := (lenNums / 4)
	var slices [4][]int
	var lens [4]int
	var i int
	for i = 0; i < 3; i++ {
		slices[i] = nums[i*quarter : (i+1)*quarter]
		lens[i] = len(slices[i])
	}
	slices[i] = nums[i*quarter:]
	lens[i] = len(slices[i])

	wg.Add(4)
	for i = 0; i < 4; i++ {
		go sort(slices[i], &wg)
	}
	wg.Wait()
	retSlice1 := make([]int, 0, lens[0]+lens[1])
	retSlice2 := make([]int, 0, lens[2]+lens[3])
	wg.Add(2)
	go merge(slices[0], slices[1], &retSlice1, &wg)
	go merge(slices[2], slices[3], &retSlice2, &wg)
	wg.Wait()
	retSlice := make([]int, 0, lenNums)
	wg.Add(1)
	merge(retSlice1, retSlice2, &retSlice, &wg)
	fmt.Println("The sorted list is: ", retSlice)

}
func sort(numSlice []int, wg *sync.WaitGroup) {
	fmt.Println("sort subroutine will sort this subarray: ", numSlice)
	lenSlice := len(numSlice)
	for i := lenSlice - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if numSlice[j] > numSlice[j+1] {
				temp := numSlice[j]
				numSlice[j] = numSlice[j+1]
				numSlice[j+1] = temp
			}
		}
	}
	(*wg).Done()
}

func merge(slice1, slice2 []int, retSlice *[]int, wg *sync.WaitGroup) {
	len1 := len(slice1)
	len2 := len(slice2)
	//retSlice := make([]int, 0, len1 + len2)
	i, j := 0, 0
	for i < len1 && j < len2 {
		if slice1[i] < slice2[j] {
			*retSlice = append(*retSlice, slice1[i])
			i++
		} else {
			*retSlice = append(*retSlice, slice2[j])
			j++
		}
	}
	for ; i < len1; i++ {
		*retSlice = append(*retSlice, slice1[i])
	}
	for ; j < len2; j++ {
		*retSlice = append(*retSlice, slice2[j])
	}
	(*wg).Done()
}
