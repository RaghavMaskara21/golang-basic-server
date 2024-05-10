package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func WsConnectionValidator(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.ToLower(r.Header.Get("Connection")) != "upgrade" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Wrong Protocol")
			return
		}

		if strings.ToLower(r.Header.Get("Upgrade")) != "websocket" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Wrong Protocol")
			return
		}

		if strings.ToLower(r.Header.Get("Sec-WebSocket-Version")) == "websocket" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Wrong Protocol")
			return
		}

		if strings.ToLower(r.Header.Get("Sec-WebSocket-Key")) == "" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Wrong Protocol")
			return
		}

		next.ServeHTTP(w, r)
	})
}
