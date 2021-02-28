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
			
			var buffer []byte
			var pType uint8 = 1
			authPacket := New(&buffer, pType)

			authPacket.Write("uint32", user.UserID)
			authPacket.Write("uint32", uint32(0))
			authPacket.Write("uint32", user.UserID)
			authPacket.Write("string", user.Username)
			authPacket.Write("uint8", user.Admin)
			authPacket.Write("uint8", user.MembershipType)

			authPacket.Send(socket)

			
		}
	
	// Chat and commands
	case 3: 
		{
			command, _ := buffer.ReadString(0)
			args, _ := buffer.ReadString(0)

			fmt.Println(command, args)
			var mBuf []byte
			var pType uint8 = 7
			
			messagePacket := New(&mBuf, pType)

			messagePacket.Write("string", "prompt")
			messagePacket.Write("string", args)

			messagePacket.Send(socket)
		}

	// Clicks and keys
	case 6:
		{
			click, _ := buffer.ReadByte()
			key, _ := buffer.ReadString(0)

			fmt.Println(click, key)
		}
	default:
		break

	}
}
