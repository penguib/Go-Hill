package buffers

import (
	"bytes"
	"fmt"
	"net"

	"Go-Hill/api"
)

var _Game Game

// HandlePacketType handles client packets
func HandlePacketType(packetType uint8, socket *net.Conn, buffer *bytes.Buffer, lPlayer *Player) {
	switch packetType {

	// Authentication
	case 1:
		{
			user, err := api.CheckAuth(socket, buffer)
			if err != nil {
				fmt.Printf("<Client: %s> Failed authentication\n", (*socket).RemoteAddr())
				break
			}

			var aBool string
			if (*user).Admin == 1 {
				aBool = "true"
			} else {
				aBool = "false"
			}

			fmt.Printf("Successfully verified! (Username: %s | ID: %s | Admin: %s)\n", (*user).Username, fmt.Sprint((*user).UserID), aBool)

			authPacket := New(&[]byte{}, Enums["Authentication"])

			authPacket.Write("uint32", user.UserID)
			authPacket.Write("uint32", uint32(0))
			authPacket.Write("uint32", user.UserID)
			authPacket.Write("string", user.Username)
			authPacket.Write("uint8", user.Admin)
			authPacket.Write("uint8", user.MembershipType)

			authPacket.Send(socket)

			newPlayer := &Player{
				Socket:     socket,
				NetID:      user.UserID,
				BrickCount: uint32(0),
				UserID:     user.UserID,
				Username:   user.Username,
				Admin:      user.Admin,
				Position:   Vector3{0, 0, 0},
				Rotation:   Vector3{0, 0, 0},
				Scale:      Vector3{1, 1, 1},
				Team:       0,
				Score:      -1,
				Game:       &_Game,
			}

			_Game.Players = append(_Game.Players, *newPlayer)
			_Game._SendClients(newPlayer)

			_Game.MessageAll("\\c6[SERVER]: \\c0" + newPlayer.Username + " has joined the server!")

			*lPlayer = *newPlayer
		}

	// Player position
	case 2:
		break

	// Chat and commands
	case 3:
		{
			command, _ := buffer.ReadString(0)
			args, _ := buffer.ReadString(0)

			if command[:len(command)-1] != "chat" {
				break
			}

			_Game.MessageAll("\\c6" + lPlayer.Username + ": \\c0" + args)

		}

	// Projectiles
	case 4:
		break

	// Brick click detection
	case 5:
		break

	// Clicks and keys
	case 6:
		{
			// click, _ := buffer.ReadByte()
			// key, _ := buffer.ReadString(0)
		}
	default:
		break

	}
}
