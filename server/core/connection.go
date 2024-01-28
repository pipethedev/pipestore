package core

import (
	"fmt"
	"net"
	"pipebase/server/types"
	"sync"
)

var sessions = make(map[net.Conn]*types.Session)
var mutex = &sync.Mutex{}

func StartTCP() {
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

		go handleAuthentication(conn)
	}
}

func handleAuthentication(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connection established from:", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Error reading: ", err)
		return
	}

	authenticationData := buffer[:n]

	fmt.Println("auth data", authenticationData)

	if !authenticate(conn, authenticationData) {
		fmt.Println("Authentication failed for:", conn.RemoteAddr())
		return
	}

	fmt.Println("Authentication successful for:", conn.RemoteAddr())

	session := &types.Session{
		Conn:     conn,
		Active:   true,
		Username: "username",
		APIKey:   "password",
	}

	mutex.Lock()
	sessions[conn] = session
	mutex.Unlock()

	go handleConnection(conn)
}

func authenticate(conn net.Conn, authData []byte) bool {
	//Todo: This could use a JWT authentication
	return true
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connection established from:", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	data := buffer[:n]

	fmt.Println("Received data", string(data))

	response := []byte("Connected to pipebase db")

	conn.Write(response)
}
