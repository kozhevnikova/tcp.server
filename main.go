package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	timeout := 3 * time.Second
	reader := bufio.NewReader(conn)

	fmt.Println("Handling new connection.")
	defer func() {
		fmt.Println("Closing connection.")
		conn.Close()
	}()

	for {
		conn.SetReadDeadline(time.Now().Add(timeout))
		text, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s", text)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3333")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		listener.Close()
		fmt.Println("Listener closed")
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}

		go handleConnection(conn)
	}
}
