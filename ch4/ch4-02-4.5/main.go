package main

import (
	"fmt"
)

/*

练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。

https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch4/ch4-02-2.md
*/

func main() {
	s := []string{"h", "e", "e", "l", "l", "o", "w", "o"}
	fmt.Printf("origin s: %v\n", s)
	s = remove(s)
	fmt.Printf("%s", s)
}

func remove(s []string) []string {
	for i := range s {
		if i > len(s)-1 {
			return s
		}
		if i < len(s)-1 && s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			return remove(s)
		}
	}
	return s

}
