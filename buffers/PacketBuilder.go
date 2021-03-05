package buffers

import (
	"Go-Hill/utils"
	"encoding/binary"
	"log"
	"math"
	"net"
)

// Enums define what type of packet you're building
var Enums = map[string]uint8{
	"Authentication":     1,
	"SendBrick":          2,
	"SendPlayers":        3,
	"Figure":             4,
	"RemovePlayer":       5,
	"Chat":               6,
	"PlayerModification": 7,
	"Kill":               8,
	"Brick":              9,
	"Team":               10,
	"Tool":               11,
	"Bot":                12,
	"ClearMap":           14,
	"DestroyBot":         15,
	"DeleteBrick":        16,
}

// PacketBuilder builds packets
type PacketBuilder struct {
	buffer []byte
}

// New creates a new packet
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
			sBuf := []byte(data.(string))
			sBuf = append(sBuf, 0)
			p.buffer = append(p.buffer, sBuf...)
		}
	case "bool":
		fallthrough
	case "uint8":
		{
			p.buffer = append(p.buffer, data.(uint8))
		}
	case "uint16":
		{
			buf := make([]byte, 2)
			binary.LittleEndian.PutUint16(buf, data.(uint16))
			p.buffer = append(p.buffer, buf...)
		}
	case "uint32":
		{
			buf := make([]byte, 4)
			binary.LittleEndian.PutUint32(buf, data.(uint32))
			p.buffer = append(p.buffer, buf...)
		}
	case "float":
		{
			buf := make([]byte, 4)
			binary.LittleEndian.PutUint32(buf, math.Float32bits(data.(float32)))
			p.buffer = append(p.buffer, buf...)
		}
	case "int32":
		{
			buf := make([]byte, 4)
			binary.LittleEndian.PutUint32(buf, uint32(data.(int32)))
			p.buffer = append(p.buffer, buf...)
		}
	default:
		break

	}

	return p.buffer
}

// Send sends the packet to the specified client
func (p *PacketBuilder) Send(socket *net.Conn) {
	utils.WriteUIntV(&p.buffer)
	(*socket).Write(p.buffer)
}

// Broadcast sends the packet to all clients
func (p *PacketBuilder) Broadcast() {
	utils.WriteUIntV(&p.buffer)

	for _, v := range _Game.Players {
		_, err := (*v.Socket).Write(p.buffer)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BroadcastExcept sends the packets to all clients except the specified
func (p *PacketBuilder) BroadcastExcept(id uint32) {
	utils.WriteUIntV(&p.buffer)

	for _, v := range _Game.Players {
		if v.NetID == id {
			continue
		}
		_, err := (*v.Socket).Write(p.buffer)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Insert inserts uint8 into the second position of the buffer
func (p *PacketBuilder) Insert(i uint8) {
	p.buffer = append(p.buffer, 0)
	copy(p.buffer[2:], p.buffer[1:])
	p.buffer[1] = i
}
