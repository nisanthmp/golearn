package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	food, locomotion, noise string
}

func (an animal) Eat() {
	fmt.Println(an.food)
}

func (an animal) Move() {
	fmt.Println(an.locomotion)
}

func (an animal) Speak() {
	fmt.Println(an.noise)
}

func main() {
	cow := animal{"grass", "walk", "moo"}
	bird := animal{"worms", "fly", "peep"}
	snake := animal{"mice", "slither", "hsss"}

	for {
		var anim, act string

		fmt.Print("> ")
		//fmt.Scan(&anim, &act)
		br := bufio.NewReader(os.Stdin)
		line, _, err := br.ReadLine()
		if err != nil {
			panic("Error reading line")
		}
		sline := strings.Trim(string(line), "\n")
		cmdString := strings.Split(sline, " ")

		anim = cmdString[0]
		act = cmdString[1]

		var an animal

		switch anim {
		case "cow":
			an = cow
		case "bird":
			an = bird
		case "snake":
			an = snake
		default:
			fmt.Println("Accepted animals are cow, bird and snake. Try again...")
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
	}
}
