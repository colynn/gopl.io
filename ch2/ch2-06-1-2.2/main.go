/*

练习 2.2： 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，
如果缺省的话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，
长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。

https://github.com/colynn/gopl-zh.github.com/blob/master/ch2/ch2-06-1.md

*/
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/ch2-06-1-2.2/lwconv"
)

//!+
func main() {

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise 2.2: %v\n", err)
			os.Exit(1)
		}
		p := lwconv.Pound(t)
		f := lwconv.Feet(t)
		fmt.Printf("%s = %s, %s = %s\n",
			p, lwconv.PToK(p), f, lwconv.FToM(f))
	}
}
