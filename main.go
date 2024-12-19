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
	str := "a"
	ans := 0

	for i, ch := range str {
		ma := make(map[string]int)
		ma[string(ch)]++
		for j := i + 1; j < len(str); j++ {
			_, ok := ma[string(str[j])]
			if ok {
				ans = max(ans, j-i)
				break
			}
		}
	}
	fmt.Println(ans)

}
