package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("tattarrattat"))
}
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	ri, rj, max := 0, 0, 0
	c := 0
	//中心扩展法
	for i := 0; i < len(s); i++ {
		//奇数情况
		for j := 0; j <= i && i+j < len(s); j++ {
			if s[i-j] != s[i+j] {
				break
			}
			c = 2*j + 1
			if c > max {
				max = c
				ri = i - j
				rj = i + j
			}
		}
		//偶数情况
		for j := 0; j <= i && i+j+1 < len(s); j++ {
			if s[i-j] != s[i+j+1] {
				break
			}
			c = 2 * j +2
			if c > max {
				max = c
				ri = i - j
				rj = i + j + 1
			}
		}
	}
	return s[ri : rj+1]
}
