/*
# Instructions

Given a word, compute the Scrabble score for that word.

## Letter Values

You'll need these:

```text
Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
```

## Examples

"cabbage" should be scored as worth 14 points:

- 3 points for C
- 1 point for A, twice
- 3 points for B, twice
- 2 points for G
- 1 point for E

And to total:

- `3 + 2*1 + 2*3 + 2 + 1`
- = `3 + 2 + 6 + 3`
- = `5 + 9`
- = 14

## Extensions

- You can play a double or a triple letter.
- You can play a double or a triple word.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Input Word :")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	fmt.Println("You entered : ", line)
	line = strings.Trim(line, "\n")
	fmt.Println("Let's get your scrabble score")
	fmt.Println("It is  :  ", Score(line))

}

func Score(word string) int {

	mymap := map[int]string{1: "A, E, I, O, U, L, N, R, S, T", 2: "D,G", 3: "B, C, M, P", 4: "F, H, V, W, Y", 5: "K", 8: "J,X", 10: "Q,Z"}
	sp := strings.Split(word, "")
	count := 0
	for _, v := range sp {
		v = strings.ToUpper(v)
		for k1, v1 := range mymap {
			if strings.Contains(v1, v) {
				count += k1
			}
		}
	}
	return count
}
