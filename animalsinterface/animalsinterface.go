package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal interface {
	Eat()
	Move()
	Speak()
}
type cow struct {
}
type bird struct {
}
type snake struct {
}

func (c cow) Eat() {
	fmt.Println("grass")
}

func (c cow) Move() {
	fmt.Println("walk")
}

func (c cow) Speak() {
	fmt.Println("moo")
}

func (c bird) Eat() {
	fmt.Println("worms")
}

func (c bird) Move() {
	fmt.Println("fly")
}

func (c bird) Speak() {
	fmt.Println("peep")
}

func (c snake) Eat() {
	fmt.Println("mice")
}

func (c snake) Move() {
	fmt.Println("slither")
}

func (c snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]animal)

	for {
		var cmd, name, anim, act string

		fmt.Print("> ")
		br := bufio.NewReader(os.Stdin)
		line, _, err := br.ReadLine()
		if err != nil {
			panic("Error reading line")
		}
		sline := strings.Trim(string(line), "\n")
		cmdString := strings.Split(sline, " ")

		cmd = cmdString[0]
		name = cmdString[1]
		if cmd == "newanimal" {
			anim = cmdString[2]
			var an animal

			switch anim {
			case "cow":
				an = cow{}
			case "bird":
				an = bird{}
			case "snake":
				an = snake{}
			default:
				fmt.Println("Accepted animals are cow, bird and snake. Try again...")
				continue
			}
			animals[name] = an
		} else if cmd == "query" {
			act = cmdString[2]
			an, found := animals[name]
			if found == false {
				fmt.Println(name, "not found. Try again...")
				continue
			}
			switch act {
			case "eat":
				an.Eat()
			case "move":
				an.Move()
			case "speak":
				an.Speak()
			default:
				fmt.Println("Accepted actions are eat, move and speak. Try again...")
			}
		} else {
			fmt.Println("Accepted commands are newanimal and query. Try again...")
		}
	}
}
