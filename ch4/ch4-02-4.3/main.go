package main

import "fmt"

/*
练习 4.3： 重写reverse函数，使用数组指针代替slice。

https://github.com/colynn/gopl-zh.github.com/blob/master/ch4/ch4-02-2.md

//!+rev  origin reverse func
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

*/

// reverse
func reverse(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func main() {
	var arr []int = []int{1, 2, 3, 6, 48, 299, 4990}
	fmt.Println("In main(), arr values: ", arr)
	reverse(&arr)
	fmt.Println("In main(), arr reverse values: ", arr)
}
