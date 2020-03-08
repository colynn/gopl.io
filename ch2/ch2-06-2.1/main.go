/*
https://github.com/colynn/gopl-zh.github.com/blob/master/ch2/ch2-06.md

练习题：tempconv ，使用示例。
*/
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/ch2-06-2.1/tempconv"
)

//!+
func main() {
	const (
		value1 tempconv.Kelvin     = 278.0
		value2 tempconv.Fahrenheit = 100.0
	)
	fmt.Printf("%s\n", tempconv.KToC(value1).String())
	fmt.Printf("%s\n", tempconv.FToC(value2).String())

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToK(c))
	}
}
