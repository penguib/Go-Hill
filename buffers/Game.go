package buffers

// Game is the main class for holding everything
type Game struct {
	Bricks  []Brick
	Players []Player

	MOTD string
}

func _Message(message string) PacketBuilder {
	p := New(&[]byte{}, Enums["Chat"])
	p.Write("string", message)

	return p
}

// MessageAll messages everyone in the server
func (g *Game) MessageAll(message string) {
	p := _Message(message)
	p.Broadcast()
}

func (g *Game) _SendClients(player *Player) {
	if len(g.Players) <= 1 {
		return
	}

	sendPlayers := New(&[]byte{}, Enums["SendPlayers"])

	sendPlayers.Write("uint8", Enums["Authentication"])
	sendPlayers.Write("uint32", player.NetID)
	sendPlayers.Write("string", player.Username)
	sendPlayers.Write("uint32", player.UserID)
	sendPlayers.Write("uint8", player.Admin)
	sendPlayers.Write("uint8", player.MembershipType)

	sendPlayers.BroadcastExcept(player.NetID)

	var count uint8 = 0
	packet := New(&[]byte{}, Enums["SendPlayers"])
	for _, v := range g.Players {
		if v.NetID == player.NetID {
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
		packet.Send(player.Socket)
	}
}

func (g *Game) _CreateFigures(player *Player) {
	playersPacket := CreatePlayerPacket(player, "ABCDEFGHIKLMNOPQUVWXYfg")
	playersPacket.BroadcastExcept(player.NetID)

	for _, v := range _Game.Players {
		if v.NetID != player.NetID {
			playerPacket := CreatePlayerPacket(&v, "ABCDEFGHIKLMNOPQUVWXYfg")
			playerPacket.Send(player.Socket)
		}
	}

	avatarPacket := CreatePlayerPacket(player, "KLMNOPQUVW")
	avatarPacket.Broadcast()
}
