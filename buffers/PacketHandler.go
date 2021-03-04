package buffers

import (
	"bytes"
	"fmt"
	"net"

	"Go-Hill/api"
	"Go-Hill/classes"
)

// Players are all the players in the game
var Players []classes.Player

// HandlePacketType handles client packets
func HandlePacketType(packetType uint8, socket *net.Conn, buffer *bytes.Buffer) {
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

			newPlayer := &classes.Player{
				Socket:     socket,
				NetID:      user.UserID,
				BrickCount: uint32(0),
				UserID:     user.UserID,
				Username:   user.Username,
				Admin:      user.Admin,
				Position:   classes.Vector3{0, 0, 0},
				Rotation:   classes.Vector3{0, 0, 0},
				Scale:      classes.Vector3{1, 1, 1},
				Team:       0,
				Score:      -1,
			}

			Players = append(Players, *newPlayer)

			if len(Players) > 1 {
				sendPlayers := New(&[]byte{}, Enums["SendPlayers"])

				sendPlayers.Write("uint8", Enums["Authentication"])
				sendPlayers.Write("uint32", newPlayer.NetID)
				sendPlayers.Write("string", newPlayer.Username)
				sendPlayers.Write("uint32", newPlayer.UserID)
				sendPlayers.Write("uint8", newPlayer.Admin)
				sendPlayers.Write("uint8", newPlayer.MembershipType)

				sendPlayers.BroadcastExcept(newPlayer.NetID)

				var count uint8 = 0
				packet := New(&[]byte{}, Enums["SendPlayers"])
				for _, v := range Players {
					if v.NetID == newPlayer.NetID {
						continue
					}
					packet.Write("uint8", Enums["Authentication"])
					packet.Write("uint32", v.NetID)
					packet.Write("string", v.Username)
					packet.Write("uint32", v.UserID)
					packet.Write("uint8", v.Admin)
					packet.Write("uint8", v.MembershipType)
				}

				count++

				if count > 0 {
					packet.Insert(count)
					packet.Send(newPlayer.Socket)
				}
			}

			playersPacket := CreatePlayerIDBuffer(newPlayer, "ABCDEFGHIKLMNOPQUVWXYfg")
			playersPacket.BroadcastExcept(newPlayer.NetID)

			for _, v := range Players {
				if v.NetID != newPlayer.NetID {
					playerPacket := CreatePlayerIDBuffer(&v, "ABCDEFGHIKLMNOPQUVWXYfg")
					playerPacket.Send(newPlayer.Socket)
				}
			}

			avatarPacket := CreatePlayerIDBuffer(newPlayer, "KLMNOPQUVW")
			avatarPacket.Broadcast()

		}

	// Player position
	case 2:
		break

	// Chat and commands
	case 3:
		{
			command, _ := buffer.ReadString(0)
			args, _ := buffer.ReadString(0)

			if command != "chat" {
				break
			}

			messagePacket := New(&[]byte{}, Enums["PlayerModification"])

			messagePacket.Write("string", "prompt")
			messagePacket.Write("string", string(args[:len(args)-1]))

			messagePacket.Broadcast()

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
