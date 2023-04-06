package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const chunkSplitSize = 4

func main() {
	fmt.Print("Enter numbers (space-separated): ")

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Error reading input:", scanner.Err())
		return
	}

	input := scanner.Text()
	if len(input) == 0 {
		fmt.Println("No numbers entered")
		return
	}

	fields := bufio.NewScanner(strings.NewReader(input))
	fields.Split(bufio.ScanWords)

	var numbers []int
	for fields.Scan() {
		num, err := strconv.Atoi(fields.Text())
		if err != nil {
			fmt.Printf("Invalid number: %s\n", fields.Text())
			return
		}
		numbers = append(numbers, num)
	}

	if err := fields.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var wg sync.WaitGroup

	wg.Add(chunkSplitSize)

	chunkSize := len(numbers) / chunkSplitSize
	for i := 0; i < chunkSplitSize; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == 3 {
			end = len(numbers)
		}
		subarray := numbers[start:end]
		fmt.Printf("Sorting subarray %d: %v\n", i+1, subarray)
		go func() {
			sort.Ints(subarray)
			wg.Done()
		}()
	}

	wg.Wait()
	sort.Ints(numbers)
	fmt.Printf("Final sorted merged list: %d\n", numbers)
}
