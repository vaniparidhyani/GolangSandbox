/*
# Instructions

Convert a trinary number, represented as a string (e.g. '102012'), to its
decimal equivalent using first principles.

The program should consider strings specifying an invalid trinary as the
value 0.

Trinary numbers contain three symbols: 0, 1, and 2.

The last place in a trinary number is the 1's place. The second to last
is the 3's place, the third to last is the 9's place, etc.

```shell
# "102012"
    1       0       2       0       1       2    # the number
1*3^5 + 0*3^4 + 2*3^3 + 0*3^2 + 1*3^1 + 2*3^0    # the value
  243 +     0 +    54 +     0 +     3 +     2 =  302
```

trinary_test.go:33: ParseTrinary("10") = 1, want 3
    trinary_test.go:33: ParseTrinary("201") = 11, want 19
    trinary_test.go:33: ParseTrinary("0201") = 33, want 19
    trinary_test.go:33: ParseTrinary("0000000000000000000000000000000000000000201") = -9223372036854775808, want 19
    trinary_test.go:33: ParseTrinary("2021110011022210012102010021220101220221") = 7794616428515451904, want 9223372036854775807
    trinary_test.go:31: ParseTrinary("2021110011022210012102010021220101220222") = -9223372036854775808, <nil>, expected error

--- FAIL: TestParseTrinary (0.00s)
    trinary_test.go:33: ParseTrinary("2021110011022210012102010021220101220221") = -9223372036854775794, want 9223372036854775807
    trinary_test.go:31: ParseTrinary("2021110011022210012102010021220101220222") = -9223372036854775793, <nil>, expected error
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Input the number:")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	fmt.Println("You entered : ", line)
	line = strings.Trim(line, "\n")
	fmt.Println("Let's convert to decimal")
	fmt.Println(ParseTrinary(line))
}

func IntPow(n, m int) int64 {
	if m == 0 {
		return 1
	}
	result := int64(n)
	for i := 2; i <= m; i++ {
		result *= int64(n)
	}
	return result
}

func ParseTrinary(s string) (int64, error) {
	sp := strings.Split(s, "")
	le := len(sp) - 1
	var count int64 = 0
	var over = errors.New("integer overflow!")
	var under = errors.New("integer underflow!")
	for k := le; k >= 0; k-- {
		v := sp[k]
		po := IntPow(3, le-k)
		val, _ := strconv.ParseInt(v, 10, 64)
		count += val * po
		if count < 0 {
			if (count) <= (math.MinInt64) {
				return 0, under
			}
		} else if (count) > 0 {
			if (count) > (math.MaxInt64) {

				return 0, over
			}
		}
	}

	return count, nil
}
