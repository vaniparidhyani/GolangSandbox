package main

import (
	"bufio"
	"fmt"
	"os"
        "strings"
	"strconv"
	"sort"
)


func main() {
        sli := make([]int, 3)
        index := 0
	fmt.Println("Enter Integers to Add to Slice. Enter 'X' when done entering.")
	for {

		inputstr := bufio.NewReader(os.Stdin)
		line, _ := inputstr.ReadString('\n')
                if strings.Contains(line, "X") {
			break
		} else {
			num, _ := strconv.Atoi(strings.TrimSpace(line))
			if index == 3 {
				sli = append(sli,num)
			} else {
				sli[index] = num
                        	index++
			}                    
		}
	}
	sort.Ints(sli)
	fmt.Println(sli)   
}
