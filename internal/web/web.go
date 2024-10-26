package web

import (
	"fmt"
	"net/http"

	"github.com/rustymotors/gorace/internal/helpers"
	"github.com/rustymotors/gorace/internal/shard"
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

		request.HandleFunc("/ShardList/", func(w http.ResponseWriter, r *http.Request) {
			shard.HandleShardList(w, r)
		})

		request.HandleFunc("/", http.NotFound)
		http.ListenAndServe(":3000", nil)
	}()
}

type AuthLoginResponse struct {
	Valid bool
	Ticket string
	ReasonCode string
	ReasonText string
	ReasonUrl string
}

func (r *AuthLoginResponse) IsValid() bool {
	return r.Valid
}

func (r *AuthLoginResponse) GetTicket() string {
	return r.Ticket
}

func (r *AuthLoginResponse) GetReasonCode() string {
	return r.ReasonCode
}

func (r *AuthLoginResponse) GetReasonText() string {
	return r.ReasonText
}

func (r *AuthLoginResponse) GetReasonUrl() string {
	return r.ReasonUrl
}

func NewAuthLoginResponse() *AuthLoginResponse {
	return &AuthLoginResponse{}
}

func (r *AuthLoginResponse) SetValid(ticket string) {
	r.Valid = true
	r.Ticket = ticket
}

func (r *AuthLoginResponse) SetInvalid(reasonCode string, reasonText string, reasonUrl string) {
	r.Valid = false
	r.ReasonCode = reasonCode
	r.ReasonText = reasonText
	r.ReasonUrl = reasonUrl
}

func (r *AuthLoginResponse) formatValidResponse() string {
	return fmt.Sprintf("Valid=TRUE\nTicket=%s", r.Ticket)
}

func (r *AuthLoginResponse) formatInvalidResponse() string {
	return fmt.Sprintf("reasoncode=%s\nreasontext=%s\nreasonurl=%s", r.ReasonCode, r.ReasonText, r.ReasonUrl)
}

func (r *AuthLoginResponse) formatResponse() string {
	if r.Valid {
		return r.formatValidResponse()
	} else {
		return r.formatInvalidResponse()
	}
}




func handleAuthentication(r *http.Request, w http.ResponseWriter) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	fmt.Println("Authenticating user: ", username)

	var userId = AuthenticateUser(username, password)

	authResponse := NewAuthLoginResponse()

	if userId > 0 {
		fmt.Println("User #", userId, " authenticated")
		authResponse.SetValid("d316cd2dd6bf870893dfbaaf17f965884e")
		helpers.WriteResponse(w, authResponse.formatResponse())
	} else {
		fmt.Println("User not authenticated")
		authResponse.SetInvalid("INV-200", "Opps~", "https://www.winehq.com")
		helpers.WriteResponse(w, authResponse.formatResponse())
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