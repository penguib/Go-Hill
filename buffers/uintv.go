package buffers

import (
	"bytes"
	"encoding/binary"
)

// Data is UIntv return type
type Data struct {
	MessageSize uint
	End         uint8
}

// ReadUIntV reads buffer determining size
func ReadUIntV(buffer *[]byte) *Data {
	reader := new(bytes.Buffer)

	if ((*buffer)[0] & 1) != 0 {
		return &Data{
			MessageSize: uint((*buffer)[0] >> 1),
			End:         1,
		}
	} else if ((*buffer)[0] & 2) != 0 {
		b := make([]byte, 2)
		uint16LE, _ := reader.Read(b)
		return &Data{
			MessageSize: uint((uint16LE >> 2) + 0x80),
			End:         2,
		}
	} else if ((*buffer)[0] & 4) != 0 {
		size := (uint16((*buffer)[2]) << 13) + uint16(((*buffer)[1]<<5)+((*buffer)[0]>>3)) + 0x4080
		return &Data{
			MessageSize: uint(size),
			End:         3,
		}
	} else {
		b := make([]byte, 4)
		uint32LE, _ := reader.Read(b)
		size := (uint32LE / 8) + 0x204080
		return &Data{
			MessageSize: uint(size),
			End:         3,
		}
	}
}

// WriteUIntV writes size of message at beginning of the buffer
func WriteUIntV(buffer *[]byte) {
	length := len(*buffer)

	if length < 0x80 {
		size := make([]byte, 1)
		size[0] = byte((length << 1) + 1)

		*buffer = append(size, (*buffer)...)
	} else if length < 0x4080 {
		size := make([]byte, 2)
		i := ((length - 0x80) << 2) + 2
		binary.LittleEndian.PutUint16(size, uint16(i))

		*buffer = append(size, (*buffer)...)
	} else if length < 0x204080 {
		size := make([]byte, 3)
		writeValue := ((length - 0x4080) << 3) + 4
		size[0] = byte(writeValue & 0xFF)
		binary.LittleEndian.PutUint16(size, uint16(writeValue>>8))

		*buffer = append(size, (*buffer)...)
	} else {
		size := make([]byte, 4)
		binary.LittleEndian.PutUint32(size, uint32((length-0x204080)*8))

		*buffer = append(size, (*buffer)...)
	}
}
