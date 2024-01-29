package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"pipebase/server/config"
	"pipebase/server/types"
	"sync"
	"sync/atomic"
)

var maxConnections = config.LoadConfig().MaxConnections

var sessions = make(map[net.Conn]*types.Session)
var mutex = &sync.Mutex{}

var connectionCount int64

func StartTCP(port int) {
	address := fmt.Sprintf(":%d", port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	defer listener.Close()

	fmt.Printf("Server started. Listening on port %s\n", address)

	connectionPool := make(chan struct{}, maxConnections)

	for {
		select {
		case connectionPool <- struct{}{}:
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				<-connectionPool
				continue
			}

			atomic.AddInt64(&connectionCount, 1)
			fmt.Printf("Connection #%d established from: %s\n", atomic.LoadInt64(&connectionCount), conn.RemoteAddr())
			go handleAuthentication(conn, connectionPool)
		default:
			fmt.Println("Connection rejected: Connection pool is full")
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				continue
			}
			conn.Close()
		}
	}
}

func handleAuthentication(conn net.Conn, connectionPool chan struct{}) {
	defer func() {
		conn.Close()
		<-connectionPool
	}()

	fmt.Println("Connection established from:", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Error reading: ", err)
		return
	}

	authenticationData := buffer[:n]

	var authStruct types.AuthRequestStruct

	err = json.Unmarshal(authenticationData, &authStruct)

	if err != nil {
		fmt.Println("Invalid authentication request:", err)
		return
	}

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

		var recordStruct types.RecordRequestStruct

		err = json.Unmarshal(data, &recordStruct)

		if err != nil {
			fmt.Println("Invalid record request:", err)
			return
		}

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
	var authStruct types.AuthRequestStruct

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
