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
	j := 1

	for j < len(str) {
		ma[string(str[i])]++
		for {
			if ma[string(str[j])] >= 1 {
				break
			}
			ma[string(str[i])] = 0
			i++
		}
		ma[s]
	}

	fmt.Println(ans)
}
