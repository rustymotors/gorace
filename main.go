package main

import (
	"fmt"
	"net"
	"github.com/rustymotors/gorace/src"

	"github.com/eiannone/keyboard"
)





func handleGamePacket(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	req := gorace.GamePacket{}
	_, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	req.Header.Type = buf[0]
	req.Header.Size = uint16(buf[1]) | uint16(buf[2])<<8
	req.Data = buf[3:]

	switch req.GetHeader() {
	case 1:
		loginPacket := gorace.GameLoginPacket{
			GamePacket: gorace.GamePacket{},
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





func startListeningOnPort(port string) {
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

func listenForKeyboardEvents(ShutdownFlag chan bool) {
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

// ========================
// Main function
// ========================

func main() {

	var ShutdownFlag = make(chan bool)

	fmt.Println("Server started")

	// Start a web server on port 3000
	gorace.StartWebServer()

	// List of ports to listen to
	ports := []string{"8226", "8227", "8228", "7003"}

	/* Listen for incoming connections on all configured ports. Do this in a goroutine.
	* Create a channel to receive keyboard events. This will be used to stop the server.
	 */
	for _, port := range ports {
		// Listen for an incoming connection. Break the loop when a signal is received.
		startListeningOnPort(port)
	}

	go listenForKeyboardEvents(ShutdownFlag)

	// Wait for the shutdown signal
	<-ShutdownFlag

	fmt.Println("Server stopped")
}
