package packets

import (
	"gocraft/protocol"
	"net"
)

type clientKnownPacksConfig struct {
	Namespace string `json:"namespace"`
	Id        string `json:"id"`
	Version   string `json:"version"`
}

func WriteSelectKnownPacks(conn net.Conn) error {
	var buffer protocol.Buffer

	config := clientKnownPacksConfig{
		Namespace: "minecraft",
		Id:        "core",
		Version:   "1.21.1",
	}

	buffer.WriteVarInt(int32(1))
	buffer.WriteVarInt(int32(len(config.Namespace)))
	buffer.WriteString(config.Namespace)
	buffer.WriteVarInt(int32(len(config.Id)))
	buffer.WriteString(config.Id)
	buffer.WriteVarInt(int32(len(config.Version)))
	buffer.WriteString(config.Version)

	return protocol.SendPacket(conn, 0x07, buffer.Bytes())
}
