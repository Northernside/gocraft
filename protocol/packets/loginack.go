package packets

import (
	"gocraft/protocol"
	"net"
)

func WriteLoginAck(conn net.Conn) error {
	var buffer protocol.Buffer
	return protocol.SendPacket(conn, 0x03, buffer.Bytes())
}
