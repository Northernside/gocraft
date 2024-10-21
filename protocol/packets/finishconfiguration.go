package packets

import (
	"gocraft/protocol"
	"net"
)

func WriteFinishConfiguration(conn net.Conn) error {
	var buffer protocol.Buffer
	return protocol.SendPacket(conn, 0x03, buffer.Bytes())
}
