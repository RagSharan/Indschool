/*
 *Read input from STDIN. Print your output to STDOUT
 *Use fmt.Scanf to read input from STDIN and fmt. Println to write output to STDOUT
 */

package even_odd

import (
	"fmt"
	"strconv"
	"strings"
)

func readInputs() {
	//concatenate integers like strings
	var length int
	var num string
	fmt.Scanln(&length)
	var numbers = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&numbers[i])
	}

	for _, value := range numbers {
		if value%2 == 0 {
			num = num + strconv.FormatInt(int64(value), 10)
		}
	}
	if len(strings.TrimSpace(num)) == 0 {
		num = "0"
	}
	fmt.Print(num)
}

func reverseString() {

}
