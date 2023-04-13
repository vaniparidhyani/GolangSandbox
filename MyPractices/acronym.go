///determine if a given word or phrase is an isogram. An isogram (also known as a "non-pattern word") is a word or phrase without a repeating letter.

package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"regexp"
)

func main(){

	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)
    	line, _ := reader.ReadString('\n')
	fmt.Println("You entered : ",line)
	line =  strings.Trim(line, "\n")	
	fmt.Println("Let's get your acronym")
	fmt.Println("It is  :  ",Abbreviate(line))
	
}

func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	split := strings.Split(s," ")
	var acro string
	var nonalpha = regexp.MustCompile(`[^a-zA-Z]+`)
	for _,v := range split {

		if strings.Contains(v,"-") {
			if len(v) > 1 {
				sp1 := strings.Split(v,"-")
				for _,v1 := range sp1 {
					sp2 := strings.Split(v1,"")
					acro += strings.ToUpper(sp2[0])
				}
			}
		} else {
			v = nonalpha.ReplaceAllString(v,"")
			sp := strings.Split(v,"")
			acro += strings.ToUpper(sp[0])
		}
	}
	return acro
}
