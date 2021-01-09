package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type person struct{ fname, lname string }

func main() {
	persons := make([]person, 0, 10)
	var fileName string
	//tempFName := make([]byte, 0, 20)
	//tempLName := make([]byte, 0, 20)
	//tempLine := make([]byte, 0, 42)

	fmt.Print("Enter the name of the file: ")
	fmt.Scan(&fileName)

	fhandle, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Failed to open file: ", fileName)
		return
	}
	defer fhandle.Close()

	buffReader := bufio.NewReader(fhandle)
	for {
		line, _, err1 := buffReader.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				fmt.Println("Error reading from file: ", fileName)
			}
			break
		}

		names := strings.Split(string(line), " ")
		p := person{names[0], names[1]}
		persons = append(persons, p)

	}
	for _, p := range persons {
		fmt.Println(p.fname, p.lname)
	}

}
