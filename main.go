package main

import (
	"bytes"
	"fmt"
	"gocraft/protocol"
	"gocraft/protocol/packets"
	"log"
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
	go func() {
		defer close(done)
		for {
			packetID, payload, err := protocol.ReadPacket(conn)
			if err != nil {
				if err.Error() == "EOF" {
					fmt.Println("Connection closed by server")
					return
				}

				fmt.Println("Failed to read packet:", err)
				return
			}

			fmt.Println("=============== Incoming packet ===============")
			fmt.Printf("ID: %d\n", packetID)
			fmt.Printf("Payload: %s\n", payload)

			switch packetID {
			case 0x03:
				// compression packet
				threshold, err := protocol.ReadVarInt(bytes.NewReader([]byte(payload)))
				if err != nil {
					fmt.Println("Failed to read threshold:", err)
					return
				}

				fmt.Printf("Threshold: %d\n", threshold)
			}

			fmt.Println("===============================================")
		}
	}()

	log.Println("connected")
	err = packets.WriteHandshake(conn, serverAddr, port, protocolVersion)
	if err != nil {
		fmt.Println("failed to send handshake:", err)
		return
	}

	log.Println("-> handshake")
	err = packets.WriteLoginStart(conn, "Northernside")
	if err != nil {
		fmt.Println("failed to send login start:", err)
		return
	}

	log.Println("-> login start")

	<-done
}
