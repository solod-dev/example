// TCP client that connects to localhost:8080
// and sends a message, then prints the response.
package main

import (
	"solod.dev/so/net"
	"solod.dev/so/os"
)

func main() {
	message := "hello"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}

	// Resolve the server address.
	raddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	// A nil laddr lets the system choose the local address.
	conn, err := net.DialTCP("tcp", nil, &raddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send a request and read the reply.
	conn.Write([]byte(message))

	var buf [256]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		panic(err)
	}
	println(string(buf[:n]))
}
