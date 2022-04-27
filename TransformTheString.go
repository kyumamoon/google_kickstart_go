package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func stringtransform(s string, f string, set map[byte]int) int64 {
	var operations int64

	for i := 0; i < len(s); i++ { // go through stringS
		a := set[s[i]]
		var minop int64 = int64(math.Abs(float64(set[f[0]] - a)))

		for i := 0; i < len(f); i++ { // go through stringF
			temp1 := int64(math.Abs(float64(set[f[i]] - a)))
			if temp1 < minop {
				minop = temp1
			}
			temp2 := int64(math.Abs(float64((26 - set[f[i]]) + a)))
			if temp2 < minop {
				minop = temp2
			}
		}

		operations += minop
	}

	return operations

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	abcset := map[byte]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9,
		'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21,
		'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26}
	// backwards; subtract by 26 and add with current number pos. Z: 26 - 26 = 0 + 1(from A) = 1 operation to get from A to Z.

	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	testcases, _ := strconv.Atoi(scanner.Text())
	outputs := make([]int64, testcases)

	for i := 0; i < testcases; i++ {

		//var stringS string
		//var stringF string
		//fmt.Scanln(&stringS)
		//fmt.Scanln(&stringF)

		scanner.Scan()
		stringS := scanner.Text()
		scanner.Scan()
		stringF := scanner.Text()

		operations := stringtransform(stringS, stringF, abcset)

		outputs[i] = operations
		//outputs[i] = fmt.Sprintf("Case #%d: %d", i+1, operations)
	}

	for index, element := range outputs {
		//fmt.Println(element)
		fmt.Printf("Case #%d: %d\n", index+1, element)
	}
}
