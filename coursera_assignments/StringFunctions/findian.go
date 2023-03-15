package main

import (
"fmt"
"bufio"
"os"
"strings"
)

func main(){
	fmt.Printf("Enter your string")
	inputstr := bufio.NewReader(os.Stdin)
	line,err := inputstr.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}else {
		lowered := strings.ToLower(line)
		lowered = strings.TrimSuffix(lowered,"\n")
		if strings.HasPrefix(lowered,"i") && strings.HasSuffix(lowered,"n") && strings.Contains(lowered,"a") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
