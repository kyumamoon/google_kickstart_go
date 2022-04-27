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
	const mC = 100100
	buf := make([]byte, mC)
	s.Buffer(buf, mC)
	s.Scan()
	T, _ := strconv.Atoi(s.Text())
	O := make([]string, T)

	testfunc := func(N, M int, D, C, dogcount *int, S string) bool {
		//cancel := true

		for i := 0; i < N; i++ {
			if S[i] == 'C' {
				if *C > 0 {
					*C--
				} else {
					break
				}
			} else if S[i] == 'D' {
				if *D > 0 {
					*D--
					*dogcount--
					*C += M
				} else {
					break
				}
			} else {
				break
			}
		}

		if *dogcount == 0 {
			return true
		} else {
			return false
		}
	}

	for i := 0; i < T; i++ {
		var (
			N        int
			D        = new(int)
			C        = new(int)
			M        int
			dogcount = new(int)
		)

		s.Scan()
		temp := strings.Fields(s.Text())
		N, _ = strconv.Atoi(temp[0])
		*D, _ = strconv.Atoi(temp[1])
		*C, _ = strconv.Atoi(temp[2])
		M, _ = strconv.Atoi(temp[3])

		s.Scan()
		S := s.Text()
		*dogcount = strings.Count(S, "D")

		if *dogcount == 0 {
			O[i] = "YES"
		} else {
			solve := testfunc(N, M, D, C, dogcount, S)
			if solve {
				O[i] = "YES"
			} else {
				O[i] = "NO"
			}
		}
	}

	for index, element := range O {
		fmt.Printf("Case #%d: %s\n", index+1, element)
	}
}
