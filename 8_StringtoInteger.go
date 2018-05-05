package main

import (
	"math"
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	for {
		if !isSpace(str[0]) {
			break
		}
		str = str[1:]
	}

	s0 := str
	if str[0] == '-' || str[0] == '+' {
		str = str[1:]
	}

	var n int32
	for _, val := range str {
		if val > '9' || val < '0' {
			break
		}
		if n > math.MaxInt32/10 || (n == math.MaxInt32/10 && val-'0' > math.MaxInt32%10 ) {
			if s0[0] == '-' {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		n = n*10 + ( val - '0')
	}
	if s0[0] == '-' {
		n = -n
	}
	return int(n)
}

func isSpace(r byte) bool {
	switch r {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}
