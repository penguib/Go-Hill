package buffers

import (
	"bytes"
	"fmt"
	"net"

	"../api"
)

// HandlePacketType handles client packets
func HandlePacketType(packetType uint8, socket *net.Conn, buffer *bytes.Buffer) {
	switch packetType {

	// Authentication
	case 1:
		{
			user, err := api.CheckAuth(socket, buffer)
			if err != nil {
				fmt.Println(err)
				fmt.Printf("<Client: %s> Failed authentication\n", (*socket).RemoteAddr())
				break
			}

			fmt.Printf("Successfully verified! (Username: %s | ID: %s | Admin: %s)\n", (*user).Username, fmt.Sprint((*user).UserID), fmt.Sprint((*user).Admin))
		}
	default:
		break

	}
}
