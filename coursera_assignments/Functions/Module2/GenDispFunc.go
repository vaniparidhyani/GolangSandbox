package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var params [3]float64
	var t     float64
	var input string
	
	fmt.Print("\U0001F481 please, enter acceleration \U0001F680, velocity \U0001F699 & displacement \U0001F4CF: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("\U0001F6A8 Error: %v\n", err)
		return
	}

	input =  strings.Trim(input, "\n")
	arr :=   strings.Split(input, " ")
	if len(arr) != 3 {
		fmt.Printf("\U0001F6A8 Error: too few values (%d), enter 3 floating values!\n", len(arr))
		return
	}

	for  i := range params[:] {
		params[i], err = strconv.ParseFloat(arr[i], 64)
		if err != nil {
			fmt.Printf("\U0001F6A8 Error: %v\n", err)
			return
		}
	}

	fmt.Print("\U0001F481 please, enter the time \U000023F1 value (float64): ")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("\U0001F6A8 Error: %v\n", err)
		return
	}

	input =  strings.Trim(input, "\n")
	t, err = strconv.ParseFloat(input, 64)

	f := GenDisplaceFn(params[0], params[1], params[2])

	fmt.Printf("\U0001F64B resulting displacement \U0001FA90 is %f m\n", f(t))
}

func GenDisplaceFn(a float64, v float64, s float64) func(t float64) float64 {
	return func(t float64) float64 { 
		return float64(1.0) / float64(2.0) * a * math.Pow(t, float64(2.0)) + v * t + s
	}
}
