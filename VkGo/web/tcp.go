package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":5000")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("cannot connect")
			conn.Close()
			continue
		}

		fmt.Println()

		bufReader := bufio.NewReader(conn)
		fmt.Println("Start reading")

		go func(conn net.Conn) {
			defer conn.Close()

			for {
				rbyte, err := bufReader.ReadByte()

				if err != nil {
					fmt.Println("cannot read")
					break
				}

				fmt.Println(string(rbyte))
			}
		}(conn)
	}

}
