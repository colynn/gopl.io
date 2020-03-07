/*

练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。

*/

package main

import (
	"fmt"
	"os"
)

//!+
func main() {
	for k, v := range os.Args[1:] {
		fmt.Printf("%d:%v\n", k, v)
	}
}

//!-
