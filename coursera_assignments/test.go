package main

import (
	"strings"
	"fmt"
//	"strconv"
)
func main() {
  s := strings.Replace("ianianian", "ni", "in", 2)
  fmt.Println(s)
}

/*
func main() {
	i, _ := strconv.Atoi("10")
	y := i * 2
	fmt.Println(y)
}*/
