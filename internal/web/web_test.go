package web

import (
	"fmt"
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
