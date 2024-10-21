package packets

import (
	"gocraft/protocol"
	"net"
)

func WriteKeepAlive(conn net.Conn, payload string) error {
	var buffer protocol.Buffer
	buffer.WriteString(payload)
	return protocol.SendPacket(conn, 0x04, buffer.Bytes())
}
