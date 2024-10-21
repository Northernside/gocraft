package states

const (
	Handshaking   = 0
	Login         = 1
	Configuration = 2
	Play          = 3
)

var CurrentState uint8 = Handshaking

var ClientBoundPackets = map[string]map[int32]string{
	"Handshaking": {
		0x00: "Client Intention",
	},
	"Login": {
		0x00: "Hello",
		0x02: "Game Profile",
		0x03: "Login Acknowledged",
	},
	"Configuration": {
		0x00: "Client Settings",
		0x02: "Custom Payload",
		0x03: "Finish Configuration",
		0x07: "Select Known Packs",
	},
}

var ServerBoundPackets = map[string]map[int32]string{
	"Handshaking": {},
	"Login": {
		0x02: "Game Profile",
		0x03: "Login Compression",
	},
	"Configuration": {
		0x01: "Custom Payload",
		0x03: "Finish Configuration",
		0x07: "Registry Data",
		0x0C: "Update Enabled Features",
		0x0D: "Update Tags",
		0x0E: "Select Known Packs",
	},
	"Play": {
		0x2B: "Login",
	},
}

func GetName(state uint8, packetId int32) string {
	if CurrentState == Handshaking {
		if _, ok := ClientBoundPackets["Handshaking"][packetId]; ok {
			return ClientBoundPackets["Handshaking"][packetId]
		}
	} else if CurrentState == Login {
		if _, ok := ClientBoundPackets["Login"][packetId]; ok {
			return ClientBoundPackets["Login"][packetId]
		}
	}

	return "Unknown"
}
