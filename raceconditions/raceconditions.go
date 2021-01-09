package main

import (
	"fmt"
	"time"
)

var sharedVar int

func main() {
	go func() {
		sharedVar = 0
		time.Sleep(20000)
		sharedVar++
	}()
	go func() {
		time.Sleep(20000)
		fmt.Println(sharedVar)
	}()
	time.Sleep(2000000)
}
