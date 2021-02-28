package api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
)

const (
	currentVersion = "0.3.0.3"
)

type player struct {
	Username       string
	UserID         uint32
	Admin          bool
	MembershipType uint8
}

type user struct {
	Token   string
	Version string
}

var playerID uint32 = 0

// CheckAuth handles the first packet from client -> server
func CheckAuth(socket *net.Conn, buffer *bytes.Buffer) (*player, error) {
	var USER user

	token, _ := buffer.ReadString(0)
	version, _ := buffer.ReadString(0)

	USER.Token = token[:len(token)-1]
	USER.Version = version[:len(version)-1]

	if USER.Version != currentVersion {
		return nil, errors.New("Outdated client")
	}

	// Mask later
	fmt.Printf("<Client: %s> Attempting authentication\n", (*socket).RemoteAddr())

	playerID++

	return &player{
		Username:       "Player" + fmt.Sprint(playerID),
		UserID:         playerID,
		Admin:          false,
		MembershipType: 1,
	}, nil
}
