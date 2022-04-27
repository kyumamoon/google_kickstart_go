package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func strokecounter(intN int, p string) int {
	pnew := p + "W"
	var strokes int

	//colors := [3]byte{'R', 'Y', 'B'}

	for c := 0; c < 3; c++ {
		var counter int
		var tempstroke int

		rfunc := func(b byte) {
			switch b {
			case 'R':
				counter++
			case 'P':
				counter++
			case 'O':
				counter++
			case 'A':
				counter++
			default:
				if counter > 0 {
					tempstroke++
					counter = 0
				}
			}

		}
		yfunc := func(b byte) {
			switch b {
			case 'Y':
				counter++
			case 'O':
				counter++
			case 'G':
				counter++
			case 'A':
				counter++
			default:
				if counter > 0 {
					tempstroke++
					counter = 0
				}
			}
		}
		bfunc := func(b byte) {
			switch b {
			case 'B':
				counter++
			case 'P':
				counter++
			case 'G':
				counter++
			case 'A':
				counter++
			default:
				if counter > 0 {
					tempstroke++
					counter = 0
				}
			}
		}

		for i := 0; i < len(pnew); i++ {

			switch c {
			case 0:
				rfunc(pnew[i])
			case 1:
				yfunc(pnew[i])
			case 2:
				bfunc(pnew[i])
			}

			/*if pnew[i] == colors[c] {
				counter++
			} else {
				if counter > 0 {
					tempstroke++
				}
			}*/

		}

		strokes += tempstroke

	}

	return strokes
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// red in purple, orange, gray
	// blue in purple, green, gray
	// yellow in orange, green, gray

	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	testcases, _ := strconv.Atoi(scanner.Text())
	outputs := make([]int, testcases)

	for i := 0; i < testcases; i++ {

		scanner.Scan()
		intN, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		stringP := scanner.Text()

		strokes := strokecounter(intN, stringP)

		//outputs[i] = fmt.Sprintf("Case #%d: %d", i+1, strokes)
		outputs[i] = strokes
	}

	for index, element := range outputs {
		fmt.Printf("Case #%d: %d\n", index+1, element)
		//fmt.Println(element)
	}
}
