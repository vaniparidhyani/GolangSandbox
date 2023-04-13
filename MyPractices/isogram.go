///determine if a given word or phrase is an isogram. An isogram (also known as a "non-pattern word") is a word or phrase without a repeating letter.

package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){

	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)
    	line, _ := reader.ReadString('\n')
	fmt.Println("You entered : ",line)
	fmt.Printf("Type is %T\n",line)
	line =  strings.Trim(line, "\n")	
	fmt.Println("Let's check for IsIsogram")
	if IsIsogram(line){
		fmt.Println("Yes It Is")
	} else {
		fmt.Println("No it's not")
	}
}

func IsIsogram(word string) bool {
	split := strings.Split(word," ")
	fmt.Println(split)
	for _,v := range split {
		sp := strings.Split(v,"")
		dup_freq := make(map[string]int)
		for i:=0 ; i <len(sp); i++ {
			item := strings.ToLower(sp[i]) 
			_,exist := dup_freq[item]
			if exist {
				dup_freq[item] += 1
			} else {
				dup_freq[item] = 1
			}
		}
		fmt.Println(dup_freq)
		for _,v := range dup_freq{
			if v > 1 {
				return false
			}
		}
	}
	return true
	
}

