package classes

import (
	"net"
)

type Player struct {
	Socket *net.Conn
	NetID uint32
	BrickCount uint32
	UserID uint32
	Username string
	Admin uint8
	MembershipType uint8
}

