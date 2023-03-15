package main

import "fmt"

func main(){
	var fltnumber float32

	fmt.Printf("Enter a floating point number")
	fmt.Scan(&fltnumber)
	fmt.Println(int(fltnumber))
}
