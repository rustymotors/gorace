package helpers

import (
	"net/http/httptest"
	"testing"
)

func TestWriteResponse(t *testing.T) {
	tests := []struct {
		name     string
		response string
	}{
		{"Empty response", ""},
		{"Non-empty response", "Hello, World!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			WriteResponse(rr, tt.response)

			if rr.Header().Get("Content-Type") != "text/plain" {
				t.Errorf("expected Content-Type to be 'text/plain', got '%s'", rr.Header().Get("Content-Type"))
			}

			if rr.Body.String() != tt.response {
				t.Errorf("expected body to be '%s', got '%s'", tt.response, rr.Body.String())
			}
		})
	}
}
