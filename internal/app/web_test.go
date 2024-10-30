package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticateUser(t *testing.T) {
	tests := []struct {
		username string
		password string
		expected int
	}{
		{"admin", "admin", 1},
		{"admin", "wrongpassword", 0},
		{"wronguser", "admin", 0},
		{"wronguser", "wrongpassword", 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("username=%s,password=%s", test.username, test.password), func(t *testing.T) {
			result := AuthenticateUser(test.username, test.password)
			if result != test.expected {
				t.Errorf("AuthenticateUser(%s, %s) = %d; want %d", test.username, test.password, result, test.expected)
			}
		})
	}
}

func TestHandleAuthentication(t *testing.T) {
	tests := []struct {
		username       string
		password       string
		expectedStatus int
		expectedBody   string
	}{
		{"admin", "admin", http.StatusOK, "Valid=TRUE\nTicket=d316cd2dd6bf870893dfbaaf17f965884e"},
		{"admin", "wrongpassword", http.StatusOK, "reasoncode=INV-200\nreasontext=Opps~\nreasonurl=https://www.winehq.com"},
		{"wronguser", "admin", http.StatusOK, "reasoncode=INV-200\nreasontext=Opps~\nreasonurl=https://www.winehq.com"},
		{"wronguser", "wrongpassword", http.StatusOK, "reasoncode=INV-200\nreasontext=Opps~\nreasonurl=https://www.winehq.com"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("username=%s,password=%s", test.username, test.password), func(t *testing.T) {
			req, err := http.NewRequest("GET", "/AuthLogin?username="+test.username+"&password="+test.password, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HandleAuthentication)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
			}

			if rr.Body.String() != test.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), test.expectedBody)
			}
		})
	}
}
