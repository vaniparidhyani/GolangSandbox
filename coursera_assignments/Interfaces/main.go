package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, motion, noise string
}

func (a Cow) Eat()   { fmt.Println("grass") }
func (a Cow) Move()  { fmt.Println("walk") }
func (a Cow) Speak() { fmt.Println("moo") }

type Bird struct {
	food, motion, noise string
}

func (a Bird) Eat()   { fmt.Println("worms") }
func (a Bird) Move()  { fmt.Println("fly") }
func (a Bird) Speak() { fmt.Println("peep") }

type Snake struct {
	food, motion, noise string
}

func (a Snake) Eat()   { fmt.Println("mice") }
func (a Snake) Move()  { fmt.Println("slither") }
func (a Snake) Speak() { fmt.Println("hiss") }

func main() {
	var m map[string]Animal
	m = make(map[string]Animal)

	for {
		var input, name, atype string
		fmt.Println("Please input either newanimal or query")
		fmt.Print(">")
		fmt.Scan(&input, &name, &atype)
		switch {
		case strings.Contains(input, "newanimal"):
			switch atype {
			case "cow":
				cow := Cow{}
				m[name] = cow
				fmt.Println("Created it!")
			case "bird":
				bird := Bird{}
				m[name] = bird
				fmt.Println("Created it!")
			case "snake":
				snake := Snake{}
				m[name] = snake
				fmt.Println("Created it!")
			default:
				fmt.Println("Correct input format:")
				fmt.Println("newanimal name cow/bird/snake")
			}

		case strings.Contains(input, "query"):
			switch atype {
			case "eat":
				m[name].Eat()
			case "move":
				m[name].Move()
			case "speak":
				m[name].Speak()
			default:
				fmt.Println("Correct input format:")
				fmt.Println("query name eat/move/speak")
			}
		default:
			fmt.Println("Correct input format:")
			fmt.Println("newanimal name cow/bird/snake")
			fmt.Println("OR")
			fmt.Println("query name eat/move/speak")
		}
	}
}
