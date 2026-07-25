// TCP server that listens on localhost:8080
// and echoes back any message it receives.
package main

import (
	"solod.dev/so/conc"
	"solod.dev/so/mem"
	"solod.dev/so/net"
)

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

	// Serve connections on a fixed set of worker threads.
	pool := conc.NewPool(mem.System, conc.PoolOptions{NumThreads: 4})
	defer pool.Free()

	println("listening on", "127.0.0.1:8080")

	// Accept connections and hand them to the pool.
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		// Move the connection to the heap so it outlives this iteration.
		connPtr := mem.Alloc[net.TCPConn](mem.System)
		*connPtr = conn

		// Blocks while every worker is busy and the queue is full,
		// which throttles accepting until a worker frees up.
		pool.Go(serve, connPtr)
	}
}

// serve reads one message from the connection, echoes it back,
// and closes the connection.
func serve(arg any) {
	conn := arg.(*net.TCPConn)
	defer mem.Free(mem.System, conn)
	defer conn.Close()

	var buf [256]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		return
	}
	conn.Write(buf[:n])
}
