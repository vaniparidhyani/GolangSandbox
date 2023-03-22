package main

import (
"fmt"
"strconv"
)

func main() {
	fmt.Println("Enter integers seperated by newline")
	var lines []int
	for {
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}
		intvar, _ := strconv.Atoi(input)
		lines = append(lines, intvar)
	}
	BubbleSort(lines)
	fmt.Println(lines)
}


func BubbleSort(sli []int){
	n := len(sli)
	for i:=0; i< n-1; i++ {
		for j:=0; j < n-i-1; j++ {
			if (sli[j] > sli[j+1]) {
				Swap(sli,j)
			}
		}
	}
}

func Swap(sli []int, i int){
	sli[i], sli[i+1] = sli[i+1], sli[i]
}
