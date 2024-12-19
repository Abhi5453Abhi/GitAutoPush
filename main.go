/**
Problem Statement:

Given a string s, find the length of the longest substring without repeating characters.

Input:    s = "abcabcbb"
Output: 3

Input:    s = "bbbbb"
Output: 1

**/

package main

import "fmt"

func main() {
	str := "abcabcc"
	ans := min(1, len(str))
	ma := make(map[string]int)
	//a -> 1

	i := 0
	j := 0

	for j < len(str) {
		for {
			if ma[string(str[j])] == 0 {
				break
			}
			//abcd
			//123
			ma[string(str[i])] = 0
			i++
		}
		ma[string(str[i])] = 1
		ans = max(ans, j-i+1)
		j++
	}
	fmt.Println(ans)
}
