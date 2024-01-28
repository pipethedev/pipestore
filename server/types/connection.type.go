package types

import "net"

type Session struct {
	Conn     net.Conn
	Active   bool
	Username string
	APIKey   string
}
