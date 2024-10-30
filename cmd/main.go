package main

import (
	"log"

	"github.com/rustymotors/gorace/internal/app"
)

// ========================
// Main function
// ========================

func main() {

	var ShutdownFlag = make(chan bool)

	log.Println("Server started")

	// Start a web server on port 3000
	web.StartWebServer()

	// List of ports to listen to
	ports := []string{"8226", "8227", "8228", "7003"}

	/* Listen for incoming connections on all configured ports. Do this in a goroutine.
	* Create a channel to receive keyboard events. This will be used to stop the server.
	 */
	for _, port := range ports {
		// Listen for an incoming connection. Break the loop when a signal is received.
		web.StartListeningOnPort(port)
	}

	go web.ListenForKeyboardEvents(ShutdownFlag)

	// Wait for the shutdown signal
	<-ShutdownFlag

	log.Println("Server stopped")
}
