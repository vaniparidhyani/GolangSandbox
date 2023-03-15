package main

import "fmt"

func main(){
	var x int = 1
	var y int
	var pointt *int // this means pointt variable is a pointer for an int data type
	pointt = &x //using & we can get the address of the variable
	y = *pointt //using * we can get the value stored at that address
	fmt.Println("The address of x is : ",pointt)
	fmt.Println("The value of x is: ",y)
}
