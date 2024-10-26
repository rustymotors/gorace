package web

import (
	"fmt"
	"net"

	"github.com/eiannone/keyboard"
)

type GamePacket struct {
	// Packet header
	Header struct {
		// Packet type
		Type uint8 // 1 byte
		// Packet size
		Size uint16 // 2 bytes
	}
	// Packet data
	Data []byte
}

func (p *GamePacket) GetHeader() uint8 {
	return p.Header.Type
}

func (p *GamePacket) GetData() []byte {
	return p.Data
}

type GameLoginPacket struct {
	GamePacket
	// Packet data
	Data struct {
		// User name
		Username string
		// Password
		Password string
	}
}

func (p *GameLoginPacket) GetUsername() string {
	return p.Data.Username
}

func (p *GameLoginPacket) GetPassword() string {
	return p.Data.Password
}

func handleGamePacket(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	req := GamePacket{}
	_, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	req.Header.Type = buf[0]
	req.Header.Size = uint16(buf[1]) | uint16(buf[2])<<8
	req.Data = buf[3:]

	switch req.GetHeader() {
	case 1:
		loginPacket := GameLoginPacket{
			GamePacket: GamePacket{},
			Data: struct {
				Username string
				Password string
			}{},
		}
		loginPacket.Data.Username = string(loginPacket.GetData()[:len(loginPacket.GetData())/2])
		loginPacket.Data.Password = string(loginPacket.GetData()[len(loginPacket.GetData())/2:])
		fmt.Println("Login packet received")
		fmt.Println("Username: ", loginPacket.GetUsername())
		fmt.Println("Password: ", loginPacket.GetPassword())
	default:
		fmt.Println("Unknown packet received")

		//Print the entire packet as a serialized hex string
		fmt.Println("Packet: ", fmt.Sprintf("%x", buf))

	}

	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}

func StartListeningOnPort(port string) {
	go func(port string) {
		ln, err := net.Listen("tcp", ":"+port)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			return
		}
		defer ln.Close()
		fmt.Println("Listening on port " + port)

		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				return
			}
			go handleGamePacket(conn)
		}
	}(port)
}

func ListenForKeyboardEvents(ShutdownFlag chan bool) {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press x to exit, h for help")
	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		if event.Rune == 'x' {
			fmt.Println("Shutting down server")
			ShutdownFlag <- true
			break
		}
		if event.Rune == 'h' {
			PrintHelp()
		}
	}
}

func PrintHelp() {
	fmt.Println("Press x to exit")
}
