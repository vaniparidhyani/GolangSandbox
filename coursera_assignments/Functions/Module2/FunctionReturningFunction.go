package main

import (
"fmt"
"math"
)

func GenDisplaceFn(a, v_o, s_o float64) func (float64) float64 {
	fn :=  func(t float64) float64 {
		var s float64
		s = (a * math.Pow(t,2))/2 + (v_o*t) + s_o
		return s
		}
	return fn
}

func main() {
	fmt.Println("Enter the values for acceleration, initial velocity, and initial displacement seperated by space")
	var a ,v, s, t float64
	fmt.Scanln(&a,&v,&s)
	fmt.Println("Enter the value for time")
	fmt.Scanln(&t)
	fn := GenDisplaceFn(a,v,s)
	fmt.Println(fn(t))
}
