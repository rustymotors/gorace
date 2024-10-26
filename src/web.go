package gorace

import (
	"fmt"
	"net/http"
)

func StartWebServer() {
	go func() {
		request := http.NewServeMux()

		// Log all inbound requests
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Received request: ", r.URL.Path)
			request.ServeHTTP(w, r)
		})


		request.HandleFunc("/AuthLogin", func(w http.ResponseWriter, r *http.Request) {
			handleAuthentication(r, w)
		})

		request.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		})
		http.ListenAndServe(":3000", nil)
	}()
}

func handleAuthentication(r *http.Request, w http.ResponseWriter) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	fmt.Println("Username: ", username)

	fmt.Println("Password: ", password)

	var userId = AuthenticateUser(username, password)

	if userId > 0 {
		fmt.Println("User #", userId, " authenticated")
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Valid=TRUE\nTicket=d316cd2dd6bf870893dfbaaf17f965884e")
	} else {
		fmt.Println("User not authenticated")
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "reasoncode=INV-200\nreasontext=Opps~\nreasonurl=https://www.winehq.com")
	}

}

func AuthenticateUser(username string, password string) (userId int) {
	userId = 0

	if username == "admin" && password == "admin" {
		fmt.Println("User authenticated")
		userId = 1
	}
	
	return userId
}