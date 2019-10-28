package main

import (
	"fmt"
	"strconv"
	"strings"
)

// HighAndLow is my solution for
// Codewars kata: https://www.codewars.com/kata/highest-and-lowest/go
func HighAndLow(in string) string {
	strFlds := strings.Fields(in)
	var highest, lowest int
	for i, str := range strFlds {
		val, _ := strconv.Atoi(string(str))
		if i == 0 {
			highest = val
			lowest = val
		}
		if val > highest {
			highest = val
		}
		if val < lowest {
			lowest = val
		}
	}
	return fmt.Sprintf("%d %d", highest, lowest)
}

func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")) // expect output to equal "42 -9"
}
