/*
# Instructions

Implement run-length encoding and decoding.

Run-length encoding (RLE) is a simple form of data compression, where runs
(consecutive data elements) are replaced by just one data value and count.

For example we can represent the original 53 characters with only 13.

```text
"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB"  ->  "12WB12W3B24WB"
```

RLE allows the original data to be perfectly reconstructed from
the compressed data, which makes it a lossless data compression.

```text
"AABCCCDEEEE"  ->  "2AB3CD4E"  ->  "AABCCCDEEEE"
```

For simplicity, you can assume that the unencoded string will only contain
the letters A through Z (either lower or upper case) and whitespace. This way
data to be encoded will never contain any numbers and numbers inside data to
be decoded always represent the count for the following character.

{
	{"", "", "empty string"},
	{"XYZ", "XYZ", "single characters only are encoded without count"},
	{"AABBBCCCC", "2A3B4C", "string with no single characters"},
	{"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB", "12WB12W3B24WB", "single characters mixed with repeated characters"},
	{"  hsqq qww  ", "2 hs2q q2w2 ", "multiple whitespace mixed in string"},
	{"aabbbcccc", "2a3b4c", "lowercase characters"},
}

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Input Word :")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	fmt.Println("You entered :", line)
	line = strings.Trim(line, "\n")
	fmt.Println("Let's get your encoded string")
	fmt.Println("It is  :", RunLengthEncode(line))
	fmt.Println("Let's decode your string again")
	fmt.Println("It is :", RunLengthDecode(RunLengthEncode(line)))
	fmt.Println("Lets see if you passed")
	if line == RunLengthDecode(RunLengthEncode(line)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("BOOOOO")
	}
}

func RunLengthEncode(input string) string {
	sp := strings.Split(input, "")
	out := ""
	le := len(sp) - 1
	c := 0

	for i := 0; i <= le; i++ {
		c = 1
		if i < le && sp[i] == sp[i+1] {

			for sp[i] == sp[i+1] {
				if i < le {
					i++
					c++
				}
				if i == le {
					break
				}
			}
			out += strconv.Itoa(c) + sp[i]
		} else {
			out += sp[i]
		}
	}
	return out
}

func RunLengthDecode(input string) string {
	sp := strings.Split(input, "")
	out := ""
	digitRegexp := regexp.MustCompile(`\d`)
	spaceRegexp := regexp.MustCompile(`\s`)
	a := 0
	if len(sp) == 0 {
		return ""
	}
	for {
		if digitRegexp.MatchString(sp[a]) {
			in := a + 1
			for {
				if digitRegexp.MatchString(sp[in]) {
					in++
				} else {
					break
				}
			}
			v1, _ := strconv.Atoi(strings.Join(sp[:in], ""))
			for i := 0; i < v1; i++ {
				if spaceRegexp.MatchString(sp[in]) {
					out += " "
				} else {
					out += sp[in]
				}
			}
			sp = sp[in+1:]
		} else if spaceRegexp.MatchString(sp[a]) {
			out += " "
			sp = sp[a+1:]
		} else {
			out += sp[a]
			sp = sp[a+1:]
		}
		if a == len(sp) {
			break
		}
	}
	return out
}
