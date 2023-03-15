package main

import (
	"fmt"
//	"os"
//	"time"
)

var (
	FC_RUNTIME_API string = "yoyo"
)
func main() {
	runtime := "bobo"
	if runtime != "" {
		fmt.Println("The Runtime API is : ", runtime)
	} else {
		fmt.Println("Error")
	}
}
