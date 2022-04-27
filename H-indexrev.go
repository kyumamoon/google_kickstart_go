package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	const mC = 3500001
	buf := make([]byte, mC)
	s.Buffer(buf, mC)
	s.Scan()
	T, _ := strconv.Atoi(s.Text())
	O := make([]([]int), T)

	hindex := func(N string, L []int) []int {
		arrays := make([]int, len(L))

		score := 1

		for i := 0; i < len(L); i++ { // loop 6 times
			//bb := false

			temp1 := L[0 : i+1]
			sort.Sort(sort.Reverse(sort.IntSlice(temp1)))

			for o := 0; o < i+1; o++ {
				/*if temp1[o] < o+1 {
					score = o
					bb = true
					break
				}*/
				if temp1[o] >= o+1 {
					score = o + 1
				}
			}

			/*if i >= 1 {
				if bb == false {
					score = i + 1
				}
			}*/

			arrays[i] = score
		}

		return arrays
	}

	for i := 0; i < T; i++ {
		s.Scan()
		N := s.Text()
		s.Scan()
		t1 := strings.Fields(s.Text())
		L := make([]int, len(t1))
		for i, e := range t1 {
			L[i], _ = strconv.Atoi(e)
		}

		solve := hindex(N, L)
		O[i] = solve
	}

	for index, element := range O {
		fmt.Printf("Case #%d: ", index+1)
		for i := 0; i < len(element); i++ {
			if i != len(element)-1 {
				fmt.Printf("%d ", element[i])
			} else {
				fmt.Printf("%d", element[i])
			}
		}
		fmt.Printf("\n")
	}
}
