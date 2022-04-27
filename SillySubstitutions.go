package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	const mC = 500001
	buf := make([]byte, mC)
	s.Buffer(buf, mC)
	s.Scan()
	T, _ := strconv.Atoi(s.Text())
	O := make([]string, T)

	subset := []string{"01", "12", "23", "34", "45", "56", "67", "78", "89", "90"}
	subset2 := []string{"2", "3", "4", "5", "6", "7", "8", "9", "0", "1"}

	for i := 0; i < T; i++ {
		s.Scan()
		N := s.Text()
		if N != "" {

		}
		s.Scan()
		S := s.Text()

		var counter int

		refreshcount := func() {
			counter = 0
			for i := 0; i < 10; i++ {
				counter += strings.Count(S, subset[i])
			}
		}

		refreshcount()

		for counter > 0 {

			for i := 0; i < 10; i++ {
				temp := strings.Count(S, subset[i])
				if temp > 0 {
					S = strings.Replace(S, subset[i], subset2[i], -1)
				}
			}

			refreshcount()

		}

		O[i] = S
	}

	for index, element := range O {
		fmt.Printf("Case #%d: %s\n", index+1, element)
	}
}
