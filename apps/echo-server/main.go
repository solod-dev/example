// TCP server that listens on localhost:8080
// and echoes back any message it receives.
package main

import "solod.dev/so/net"

func main() {
	// Resolve the local address to listen on.
	laddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	// Start listening on the local address.
	ln, err := net.ListenTCP("tcp", &laddr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	println("listening on", "127.0.0.1:8080")

	// Accept connections and serve them in a loop.
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		serve(&conn)
	}
}

// serve reads one message from the connection, echoes it back,
// and closes the connection.
func serve(conn *net.TCPConn) {
	defer conn.Close()

	var buf [256]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		return
	}
	conn.Write(buf[:n])
}
