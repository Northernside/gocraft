package packets

import (
	"gocraft/protocol"
	"net"
)

func WriteHandshake(conn net.Conn, serverAddr string, port int, protocolVersion int32) error {
	var buffer protocol.Buffer

	buffer.WriteVarInt(protocolVersion)
	buffer.WriteVarInt(int32(len(serverAddr)))
	buffer.WriteString(serverAddr)
	buffer.WriteUint16(uint16(port))
	buffer.WriteVarInt(2)

	return protocol.SendPacket(conn, 0x00, buffer.Bytes())
}
