/*


练习 8.8： 使用select来改造8.3节中的echo服务器，为其增加超时，这样服务器可以在客户端10秒中没有任何喊话时自动断开连接。

https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch8/ch8-07.md


*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	message := make(chan string)
	go func() {
		for input.Scan() {
			message <- input.Text()
		}
	}()

	for {
		// NOTE: ignoring potential errors from input.Err()
		select {
		case <-time.After(10 * time.Second):
			log.Printf("Client %s automatically disconnects without any shouting for 10 seconds", c.RemoteAddr())
			c.Close()
			return
		case msg := <-message:
			go echo(c, msg, 1*time.Second)
		}
	}

}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		go handleConn(conn)

	}
}
