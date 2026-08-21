// The `so/net` package provides the tools we need
// to easily build TCP socket servers.
//
// After you run this server, you can connect to it using `telnet`:
//
//	telnet localhost 8090
//
// Type a line of text and press Enter. The server will respond with
// the same text in uppercase, prefixed with "ACK:".
package main

import (
	"solod.dev/so/bufio"
	"solod.dev/so/conc"
	"solod.dev/so/fmt"
	"solod.dev/so/mem"
	"solod.dev/so/net"
	"solod.dev/so/strings"
)

func main() {
	// Resolve the local address to listen on (port 8090 on all interfaces).
	laddr, err := net.ResolveTCPAddr("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	// net.ListenTCP starts the server on the TCP network and given address.
	listener, err := net.ListenTCP("tcp", &laddr)
	if err != nil {
		panic(err)
	}
	// Close the listener to free the port when the application exits.
	defer listener.Close()

	buf := make([]byte, 1024)
	println("Server listening on", laddr.String(buf))

	// Loop indefinitely to accept new client connections.
	for {
		// Wait for a connection.
		conn, err := listener.Accept()
		if err != nil {
			println("Error accepting conn:", err)
			continue
		}

		raddr := conn.RemoteAddr()
		fmt.Printf("Accepted connection from %s:%d\n", raddr.IP.String(buf), raddr.Port)

		// Move connection to the heap so it can be safely used in a new thread.
		connPtr := mem.Alloc[net.TCPConn](mem.System)
		*connPtr = conn

		// We launch a new thread to handle the connection so that
		// the main loop can continue accepting more connections.
		// In real-world applications, prefer using conc.Pool
		// instead of spawning a new thread for each connection.
		th := conc.Go(handleConnection, connPtr)
		th.Detach()
	}
}

// handleConnection handles a single client connection,
// reading one line of text from the client and returning a response.
func handleConnection(arg any) any {
	conn := arg.(*net.TCPConn)
	defer mem.Free(mem.System, conn)

	// Closing the connection releases resources when
	// we are finished interacting with the client.
	defer conn.Close()

	// Use bufio.NewReader to read one line of data
	// from the client (terminated by a newline).
	reader := bufio.NewReader(mem.System, conn)
	defer reader.Free()
	message, err := reader.ReadString('\n')
	if err != nil {
		println("Read error:", err)
		return err
	}

	// Create and send a response back to the client,
	// demonstrating two-way communication.
	ackMsg := strings.ToUpper(mem.System, strings.TrimSpace(message))
	defer mem.FreeString(mem.System, ackMsg)
	buf := make([]byte, 1024)
	response := fmt.Sprintf(buf, "ACK: %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		println("Server write error:", err)
		return err
	}
	return nil
}
