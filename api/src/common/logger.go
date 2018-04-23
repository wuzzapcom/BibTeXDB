package common

import (
	"log"
	"net/http"
)

// LogRequest ..
func LogRequest(r http.Request) {
	log.Printf("Request: %s\nMethod: %s\nFrom: %s\nWith headers: %s\n", r.URL, r.Method, r.RemoteAddr, r.Header)
}

// LogErrorResponseWriter ..
func LogErrorResponseWriter(code int, message string) {
	log.Printf("Return error(%d): %s\n------------\n", code, message)
}
