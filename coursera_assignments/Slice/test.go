package main

import (
"fmt"
"os"
"strings"
"bufio"
"strconv"
)

func main() {
 for {
             consoleReader := bufio.NewReader(os.Stdin)
             fmt.Print(">")

             input, _ := consoleReader.ReadString('\n')

             input = strings.ToLower(input)

             if strings.HasPrefix(input, "bye") {
                fmt.Println("Good bye!")
                os.Exit(0)
             }
             num, _ := strconv.Atoi(input)
             fmt.Println(num)
 }
}
