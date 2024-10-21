package packets

import (
	"gocraft/protocol"
	"net"
)

func WriteBrand(conn net.Conn, brand string) error {
	var buffer protocol.Buffer

	buffer.WriteVarInt(int32(len("minecraft:brand")))
	buffer.WriteString("minecraft:brand")

	buffer.WriteVarInt(int32(len(brand)))
	buffer.WriteString(brand)

	return protocol.SendPacket(conn, 0x02, buffer.Bytes())
}
