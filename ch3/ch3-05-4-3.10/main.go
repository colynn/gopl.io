package main

import (
	"bytes"
	"fmt"
)

/*
练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。

https://github.com/colynn/gopl-zh.github.com/blob/master/ch3/ch3-05-4.md

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

*/

func comma(s string) string {
	var buf bytes.Buffer
	l := len(s)

	if l <= 3 {
		return s
	}
	for i := 0; i < l; i++ {
		if (l-i)%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("12345678923"))
}
