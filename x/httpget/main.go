package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "golab.io:80")
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf(status)
}
