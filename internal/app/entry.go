package web

import (
	"fmt"
	"log"
	"net"

	"github.com/eiannone/keyboard"
	"github.com/rustymotors/gorace/internal/packets"
)

func handleGamePacket(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	req := packets.GamePacket{}
	_, err := conn.Read(buf)

	if err != nil {
		log.Println("Error reading:", err.Error())
	}

	req.FromBytes(buf)

	log.Println("Message ID: ", req.MessageId())
	//Print the entire packet as a serialized hex string
	log.Println("Packet: ", fmt.Sprintf("%x", buf))

	switch req.MessageId() {
	case 261: // Login packet
		loginPacket := packets.GameLoginPacket{}
		loginPacket.FromGamePacket(req)
		log.Println("Login packet received")
		log.Println(loginPacket.ToString())
	default:
		log.Println("Unknown packet received")

		//Print the entire packet as a serialized hex string
		log.Println("Packet: ", fmt.Sprintf("%x", buf))

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
			log.Println("Error listening:", err.Error())
			return
		}
		defer ln.Close()
		log.Println("Listening on port " + port)

		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println("Error accepting: ", err.Error())
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
			log.Println("Shutdown requested by console")
			log.Println("Shutting down")
			fmt.Println("Shutting down")
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
