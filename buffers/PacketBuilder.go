package buffers

import (
	"net"
	"encoding/binary"
	"fmt"
)

const (
	Authentication = iota + 1

    SendBrick

    SendPlayers

    Figure

    RemovePlayer

    Chat

    PlayerModification

    Kill

    Brick

    Team

    Tool

    Bot

    ClearMap

    DestroyBot

    DeleteBrick
)

type PacketBuilder struct {
	buffer []byte
}


func New(buffer *[]byte, packetType uint8) PacketBuilder {
	var p PacketBuilder
	p.buffer = *buffer
	p.buffer = append(p.buffer, packetType)
	return p
}

func (p *PacketBuilder) Write(dataType string, data interface{}) []byte {
	switch dataType {
	case "string":
		{
			p.buffer = append(p.buffer, []byte(data.(string))...)
		}
	case "bool":
			fallthrough
	case "uint8":
		{
			p.buffer = append(p.buffer, data.(uint8))
		}
	case "uint16":
		{
			d := make([]byte, 2)
			binary.LittleEndian.PutUint16(d, data.(uint16))
			p.buffer = append(p.buffer, d...)
		}
	case "uint32":
		{
			d := make([]byte, 4)
			binary.LittleEndian.PutUint32(d, data.(uint32))
			p.buffer = append(p.buffer, d...)
		}
	// case "int32":
	// 	{
	// 		d := make([]byte, 4)
	// 		binary.LittleEndian.PutInt32(d, data.(int32))
	// 		p.buffer = append(p.buffer, d...)
	// 	}
	default:
		break
		
	}

	return p.buffer
}

func (p *PacketBuilder) Send(socket *net.Conn) {
	WriteUIntV(&p.buffer)
	n, err := (*socket).Write(p.buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}