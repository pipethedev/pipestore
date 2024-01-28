package utils

import "net"

func AvailablePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")

	if err != nil {
		return 0, err
	}

	defer listener.Close()

	address := listener.Addr().(*net.TCPAddr)

	return address.Port, nil
}
