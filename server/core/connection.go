package core

import (
	"fmt"
	"net"
)

func StartTCP() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server started. Listening on port :8080")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connection established from:", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Error reading: ", err)
		return
	}

	data := buffer[n]

	fmt.Println("Received data", string(data))

	response := []byte("Connected to pipebase db")

	conn.Write(response)
}
