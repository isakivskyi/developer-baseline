// This program implements a simple TCP server that listens on a specified port

// Usage: go run tcp_server.go <port> <echo_prefix>
// Example: go run -v networking/tcp_server.go 9000 hello
// In another terminal: echo python | nc localhost 9000

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	port := fmt.Sprintf(":%s", os.Args[1])
	echo_prefix := os.Args[2]

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	fmt.Printf("listening on %s, echo prefix: %s\n", listener.Addr(), echo_prefix)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}

		go handleConnection(conn, echo_prefix)
	}
}

func handleConnection(conn net.Conn, prefix string) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("request: %s", bytes)

		line := fmt.Sprintf("%s %s", prefix, bytes)
		fmt.Printf("response: %s", line)
		conn.Write([]byte(line))
	}
}
