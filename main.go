package main

import (
	"fmt"
	"gocraft/protocol"
	"gocraft/protocol/packets"
	"gocraft/protocol/states"
	"net"
)

const (
	serverAddr      = "localhost"
	port            = 25565
	protocolVersion = int32(767)
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverAddr, port))
	if err != nil {
		fmt.Println("failed to connect:", err)
		return
	}
	defer conn.Close()

	done := make(chan bool)

	packets.WriteHandshake(conn, serverAddr, port, protocolVersion)
	states.CurrentState = states.Login
	packets.WriteLoginStart(conn, "Northernside")

	go func() {
		defer close(done)
		for {
			packetId, _, _, err := protocol.ReadPacket(conn)
			if err != nil {
				if err.Error() == "EOF" {
					fmt.Println("Connection closed by server")
					return
				}

				fmt.Println("Failed to read packet:", err)
				return
			}

			switch packetId {
			case 0x02:
				if states.CurrentState == states.Login {
					fmt.Println("Login completed!")
					packets.WriteLoginAck(conn)

					states.CurrentState = states.Configuration
					packets.WriteBrand(conn, "labymod")
					packets.WriteClientSettings(conn, "en_US", 8, 0, true, 0x7F, 1, false, true)
				}
			case 0x01:
				if states.CurrentState == states.Configuration {
					fmt.Println("Server sent custom payload")
				}
			case 0x03:
				if states.CurrentState == states.Configuration {
					fmt.Println("Server sent finish configuration")
					packets.WriteFinishConfiguration(conn)
					states.CurrentState = states.Play
				}
			case 0x07:
				if states.CurrentState == states.Configuration {
					fmt.Println("Server sent registry data")
				}
			case 0x0C:
				if states.CurrentState == states.Configuration {
					fmt.Println("Server sent update enabled features")
				}
			case 0x0D:
				if states.CurrentState == states.Configuration {
					fmt.Println("Server sent update tags")
				}
			case 0x0E:
				if states.CurrentState == states.Configuration {
					fmt.Println("Server sent select known packs")
					packets.WriteSelectKnownPacks(conn)
				}
			case 0x2B:
				if states.CurrentState == states.Play {
					fmt.Println("Server sent login")
				}
			}
		}
	}()

	<-done
}
