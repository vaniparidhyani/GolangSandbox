package main

import (
"fmt"
"os"
"bufio"
"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {

    var fi string
    fmt.Printf("Enter the filename")
    fmt.Scan(&fi)

    readFile, err := os.Open(fi)

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var names []Name
    for fileScanner.Scan() {
	line := fileScanner.Text()
	splitline := strings.Fields(line)
        if ((len(splitline[0]) < 20) || (len(splitline[1]) < 20)) {
		fmt.Println("Each field will be a string of size 20. Please make sure the first and last names are of at least 20 characters")
                fmt.Println("This name length does not match the question criteria : ",splitline)
                continue
        } else {
                names = append(names,Name{fname : splitline[0][:20],lname : splitline[1][:20]})
        }
     }

     for _,value := range names {
         fmt.Printf("FirstName : %s",value.fname)
	 fmt.Println(" Last Name : ",value.lname)
     }
}
