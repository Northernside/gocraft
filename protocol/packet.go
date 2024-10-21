package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

func SendPacket(conn net.Conn, packetID int32, payload []byte) error {
	var buffer Buffer

	packetLength := VarIntSize(packetID) + len(payload)
	buffer.WriteVarInt(int32(packetLength))
	buffer.WriteVarInt(packetID)
	buffer.Write(payload)

	fmt.Printf("sending packet: ID=%d, Length=%d, Payload=%s\n", packetID, packetLength, string(payload))

	_, err := conn.Write(buffer.Bytes())
	return err
}

func ReadPacket(conn net.Conn) (int32, string, error) {
	packetLength, err := ReadVarInt(conn)
	if err != nil {
		return 0, "", err
	}

	log.Printf("packet length: %d\n", packetLength)

	packetID, err := ReadVarInt(conn)
	if err != nil {
		log.Printf("packet id read error: %s\n", err)
		return 0, "", err
	}

	log.Printf("packet ID: %d\n", packetID)

	// remaining payload
	payload := make([]byte, int(packetLength)-VarIntSize(packetID))
	_, err = conn.Read(payload)
	if err != nil {
		return 0, "", err
	}

	decodedPayload := string(payload)
	return packetID, decodedPayload, nil
}

func ReadVarInt(reader io.Reader) (int32, error) {
	var result int32 = 0
	var numRead int32 = 0
	for {
		byteValue := make([]byte, 1)
		_, err := reader.Read(byteValue)
		if err != nil {
			return 0, err
		}

		value := byteValue[0]
		result |= (int32(value) & 0x7F) << (7 * numRead)

		numRead++
		if numRead > 5 {
			return 0, fmt.Errorf("VarInt is too big")
		}

		if (value & 0x80) == 0 {
			break
		}
	}

	return result, nil
}

func VarIntSize(value int32) int {
	size := 0
	for {
		value >>= 7
		size++
		if value == 0 {
			break
		}
	}

	return size
}

func (b *Buffer) WriteVarInt(value int32) {
	for {
		temp := byte(value & 0x7F)
		value >>= 7
		if value != 0 {
			temp |= 0x80
		}

		b.WriteByte(temp)
		if value == 0 {
			break
		}
	}
}

type Buffer struct {
	bytes.Buffer
}

func (b *Buffer) WriteUint16(v uint16) error {
	return binary.Write(&b.Buffer, binary.BigEndian, v)
}

func (b *Buffer) WriteUint32(v uint32) error {
	return binary.Write(&b.Buffer, binary.BigEndian, v)
}

func (b *Buffer) WriteUint64(v uint64) error {
	return binary.Write(&b.Buffer, binary.BigEndian, v)
}
