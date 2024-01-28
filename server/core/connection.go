package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"pipebase/server/types"
	"sync"
)

var sessions = make(map[net.Conn]*types.Session)
var mutex = &sync.Mutex{}

func StartTCP(port int) {
	address := fmt.Sprintf(":%d", port)

	listener, err := net.Listen("tcp", address)

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
	fmt.Println("Connection established from:", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Error reading: ", err)
		return
	}

	authenticationData := buffer[:n]

	username, apiKey, err := extractAuthenticationCredentials(authenticationData)

	if err != nil {
		fmt.Println("Error extracting credentials:", err)
		return
	}

	if !authenticate(username, apiKey) {
		fmt.Println("Authentication failed for:", conn.RemoteAddr())
		return
	}

	fmt.Println("Authentication successful for:", conn.RemoteAddr())

	session := &types.Session{
		Conn:     conn,
		Active:   true,
		Username: username,
		APIKey:   apiKey,
	}

	mutex.Lock()
	sessions[conn] = session
	mutex.Unlock()

	go handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connection established from:", conn.RemoteAddr())

	for {
		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client:", conn.RemoteAddr())
				break
			}
			fmt.Println("Error reading:", err)
			return
		}

		data := buffer[:n]

		fmt.Println("Received data", string(data))

		response := []byte("Connected to pipebase db")

		_, err = conn.Write(response)
		if err != nil {
			fmt.Println("Error writing response:", err)
			return
		}
	}
}

func extractAuthenticationCredentials(authData []byte) (string, string, error) {
	var authStruct types.RequestStruct

	err := json.Unmarshal(authData, &authStruct)
	if err != nil {
		return "", "", err
	}

	return authStruct.Auth.Username, authStruct.Auth.APIKey, nil
}

func authenticate(userName string, apiKey string) bool {
	//Todo: This could use a JWT authentication
	return true
}
