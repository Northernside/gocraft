package packets

import (
	"encoding/binary"
	"encoding/hex"
	"gocraft/protocol"
	"net"
)

func WriteLoginStart(conn net.Conn, username string) error {
	var buffer protocol.Buffer

	buffer.WriteVarInt(int32(len(username)))
	buffer.WriteString(username)

	uuid := "97cedeb90f78362d9ed3ebb40182b5d5"
	uuidBytes, err := hex.DecodeString(uuid)
	if err != nil {
		return err
	}

	msb := binary.BigEndian.Uint64(uuidBytes[:8])
	lsb := binary.BigEndian.Uint64(uuidBytes[8:])

	buffer.WriteUint64(msb)
	buffer.WriteUint64(lsb)

	return protocol.SendPacket(conn, 0x00, buffer.Bytes())
}
