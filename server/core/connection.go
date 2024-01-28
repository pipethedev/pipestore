package core

import (
	"fmt"
	"net"
)

func StartTCP(username string, apiKey string) {
	if !authenticate(username, apiKey) {
		fmt.Println("Invalid pipestore credentials, unable to perform connection")
		return
	}

	listener, err := net.Listen("tcp", ":5771")

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server started. Listening on port :5771")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
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

func authenticate(username string, apiKey string) bool {
	return username == "pipethedev" && apiKey == "123"
}
