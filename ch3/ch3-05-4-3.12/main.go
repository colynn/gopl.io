package main

import (
	"fmt"
	"strings"
)

/*
练习 3.12： 编写一个函数，判断两个字符串是否是相互打乱的，
也就是说它们有着相同的字符，但是对应不同的顺序。

https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch3/ch3-05-4.md

*/

func main() {
	fmt.Println(verifyStr("abbce##", "abbce##"))
	fmt.Println(verifyStr("abcde", "abdce"))
	fmt.Println(verifyStr("abcded", "abcde"))
}

func strContain(str1, str2 string) bool {
	for _, v := range str1 {
		if !strings.Contains(str2, string(v)) {
			return false
		}
	}
	return true
}

func verifyStr(str1, str2 string) bool {
	if len(str1) != len(str2) {
		fmt.Printf("str length is not equal, not match: ")
		return false
	}

	if strContain(str1, str2) && strContain(str2, str1) {
		for i := range str1 {
			if str1[i] != str2[i] {
				return true
			}
		}
		fmt.Printf("It has the same characters but in the same order: ")
	}
	return false
}
