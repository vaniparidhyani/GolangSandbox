package main

import (
"fmt"
)

type Animal struct {
	food , motion , noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.motion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{"grass", "walk","moo"}
	bird := Animal{"worms","fly","peep"}
	snake := Animal{"mice","slither","hsss"}

	fmt.Println("Hi! What would you like to know about the animals? Enter in this format >cow speak")
	for {
		fmt.Print(">")
		var animal, detail string
		fmt.Scanln(&animal,&detail)
		if animal == "cow" {
			if detail == "eat" {
				cow.Eat()
			} else if detail == "move" {
				cow.Move()
			} else if detail == "speak" {
				cow.Speak()
			} else {
				fmt.Println("We can only give you details about eat, move or speak!")
			}
		} else if animal == "bird" {
                        if detail == "eat" {
                                bird.Eat()
                        } else if detail == "move" {
                                bird.Move()
                        } else if detail == "speak" {
                                bird.Speak()
                        } else {
                                fmt.Println("We can only give you details about eat, move or speak!")
                        }
                } else if animal == "snake" {
                        if detail == "eat" {
                                snake.Eat()
                        } else if detail == "move" {
                                snake.Move()
                        } else if detail == "speak" {
                                snake.Speak()
                        } else {
                                fmt.Println("We can only give you details about eat, move or speak!")
                        }
                } else {
			fmt.Println("We only have data on cow, bird and snake!")
		}
	}
}	

