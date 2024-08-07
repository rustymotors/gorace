package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Print the request path
		fmt.Println(r.URL.Path)

		// Print the request headers
		for name, headers := range r.Header {
			for _, h := range headers {
				fmt.Printf("%v: %v\n", name, h)
			}
		}

		// Print the request body as hex
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading body: %s\n", err)

			// Send an error response
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			fmt.Printf("Body: %x\n", body)
		}

		// Send a response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("rustyserver")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
		} else {
			fmt.Fprintf(w, "Output: %s", out)
		}
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			fmt.Fprintf(w, "%s: %s\n", pair[0], pair[1])
		}
	})

	http.ListenAndServe(":3000", nil)
}
