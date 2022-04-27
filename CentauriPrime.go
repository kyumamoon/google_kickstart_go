package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Answer struct {
	K string
	Y string
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	const mC = 100001
	buf := make([]byte, mC)
	s.Buffer(buf, mC)
	s.Scan()
	T, _ := strconv.Atoi(s.Text())
	O := make([]Answer, T)

	vowelset := [10]byte{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}

	who := func(N string) Answer {
		check := N[len(N)-1]

		if check == 'y' || check == 'Y' {
			return Answer{K: N, Y: "nobody"}
		} else {
			b := true
			var A Answer
			if b {
				for _, e := range vowelset {
					if check == e {
						A = Answer{K: N, Y: "Alice"}
						b = false
					}
				}
			}
			if b {
				A = Answer{K: N, Y: "Bob"}
			}

			return A
		}
	}

	for i := 0; i < T; i++ {
		s.Scan()
		N := s.Text()

		solve := who(N)

		O[i] = solve
	}

	for index, element := range O {
		fmt.Printf("Case #%d: %s is ruled by %s.\n", index+1, element.K, element.Y)
	}
}
