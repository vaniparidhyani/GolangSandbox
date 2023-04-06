package main

import (
	"bufio"
	"fmt"
	"os"
        "strings"
	"strconv"
	"sort"
	"math"
	"log"
)

// This is the goroutine which sorts the subarray. It takes an int array as input and a []int channel to save the result.
func sorter_goroutine(a [] int, c chan [] int) {
	fmt.Printf("Received this array in sorter goroutine. This array will be now sorted.   :")
	fmt.Println(a)
	sort.Ints(a) // sort package has been used here
	c <- a
}

// This function is used to split the input array into 2 equal subarrays.
func splitter(a [] int)([] int, [] int) {
	mid := int(math.Ceil(float64(len(a))/2.0))
	return a[:mid], a[mid:]
}

// This function takes user input from a single line and returns an int array.
func inputter()[] int {
	fmt.Printf("Please enter a list of integers (space-separated): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	var numlist = make([]int, 0)
	var strVal string = scanner.Text()
	strlist := strings.Split(strVal, " ")
	for _,v := range strlist {
		val, _ := strconv.Atoi(v)
		numlist = append(numlist,val)	
	}
	return numlist
}

// This is the main goroutine.
func main() {

	// Let's take the user input first
	sli := inputter()
	lent := len(sli)

	// Check to make sure we meet minimum no of integers to sort.
	if lent < 4 {
		fmt.Println("At least enter 4 elements so I can have at least 4 subarrays")
	} else {

		// Splitting the input array into 2 halves
		first,second := splitter(sli)
		slice_of_slices := make([][]int , 4)

		// Splitting again to get 4 subarrays
		slice_of_slices[0],slice_of_slices[1] = splitter(first)
		slice_of_slices[2],slice_of_slices[3] = splitter(second)

		// Initializing a channel for goroutine to pass values
		c := make(chan [] int)
		final_list := []int{}

		// For loop to call the sorter_goroutine 4 times for 4 subarrays
		for i:=0; i < 4; i++ {
	
			go sorter_goroutine(slice_of_slices[i],c)
			// Saving the sorted subarrays into a final_list
			final_list = append(final_list, <- c...)
		}

		// Let's call the sorter_goroutine once more to sort the merged final_list
		go sorter_goroutine(final_list,c)
		
		// Printing the channel to print the entire sorted list
		fmt.Println(<- c)
	}	
}
