package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"Go-Hill/buffers"
	"Go-Hill/lualang"
)

func handleConnection(c net.Conn) {
	fmt.Printf("<New client: %s\n", c.RemoteAddr())
	for {
		buf := make([]byte, 40)
		_, err := c.Read(buf)
		if err != nil {
			break
		}

		b := buffers.ReadUIntV(&buf)

		packet := buf[b.End:]
		parsedPacket := packet[:b.MessageSize]
		r := bytes.NewReader(parsedPacket)
		var packetType uint8
		var buffer *bytes.Buffer

		z, err := zlib.NewReader(r)
		if err != nil {
			buffer = bytes.NewBuffer(parsedPacket)
			packetType, _ = buffer.ReadByte()
		} else {
			p, _ := ioutil.ReadAll(z)
			buffer = bytes.NewBuffer(p)
			packetType, _ = buffer.ReadByte()
		}

		buffers.HandlePacketType(packetType, &c, buffer)

	}
	fmt.Printf("<Client: %s> Lost connection.\n", c.RemoteAddr())
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fmt.Println("Starting server...")
	go lualang.Init()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
