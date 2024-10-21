package packets

import (
	"gocraft/protocol"
	"net"
)

func WriteClientSettings(conn net.Conn, locale string, viewDistance byte, chatMode byte, chatColors bool, displayedSkinParts byte, mainHand byte, textFiltering, serverListing bool) error {
	var buffer protocol.Buffer

	buffer.WriteVarInt(int32(len(locale)))
	buffer.WriteString(locale)
	buffer.WriteUint8(viewDistance)
	buffer.WriteUint8(chatMode)
	buffer.WriteBool(chatColors)
	buffer.WriteUint8(displayedSkinParts)
	buffer.WriteUint8(mainHand)
	buffer.WriteBool(textFiltering)
	buffer.WriteBool(serverListing)

	return protocol.SendPacket(conn, 0x00, buffer.Bytes())
}
