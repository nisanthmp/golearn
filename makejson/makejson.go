package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	person := make(map[string]string)

	br := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a name: ")
	tempString, err := br.ReadString('\n')

	if err != nil {
		fmt.Println("Failed to read the name string")
		return
	}
	tempString = strings.Trim(tempString, "\n")
	person["name"] = tempString

	fmt.Println("Enter the address in a single line: ")
	tempString, err = br.ReadString('\n')

	if err != nil {
		fmt.Println("Failed to read the address string")
		return
	}
	tempString = strings.Trim(tempString, "\n")
	person["address"] = tempString

	// jsonObject is of type []byte
	jsonObject, err2 := json.Marshal(person)
	if err2 != nil {
		fmt.Println("Error marshalling map to json object")
		return
	}
	fmt.Print("JSON Object: ")
	fmt.Println(string(jsonObject))
}
