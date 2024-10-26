package main

import (
	"fmt"
	"github.com/rustymotors/gorace/src"


)







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
		gorace.StartListeningOnPort(port)
	}

	go gorace.ListenForKeyboardEvents(ShutdownFlag)

	// Wait for the shutdown signal
	<-ShutdownFlag

	fmt.Println("Server stopped")
}
