package helpers

import (
	"fmt"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, response)
}