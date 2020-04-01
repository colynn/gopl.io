package main

import (
	"bytes"
	"fmt"
	"strings"
)

/*
练习 3.11： 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。

https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch3/ch3-05-4.md

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

	var buffer bytes.Buffer
	var flag byte
	if s[0] == '-' || s[0] == '+' {
		buffer.WriteByte(s[0])
		flag = s[0]
		s = s[1:]
	}

	// 分离整数部分与小数部位
	var fractional, integerPart string
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		fractional = s[dot:]
		integerPart = s[:dot]
	} else {
		integerPart = s
	}

	num := len(integerPart)
	var buf bytes.Buffer

	if flag != 0 {
		buf.WriteByte(flag)
	}
	for i := 0; i < num; i++ {
		if (num-i)%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(integerPart[i])
	}

	// 添加小数点后的部分
	buf.WriteString(fractional)
	return buf.String()
}

func main() {
	fmt.Println(comma("+1231112.3456"))
	fmt.Println(comma("-1234567890.3456"))
}
