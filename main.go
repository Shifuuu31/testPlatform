package main

import (
	"fmt"
	"net/http"
	"strings"
)

func detectPlatform(userAgent string) string {
	if strings.Contains(userAgent, "Android") {
		return "Android"
	} else if strings.Contains(userAgent, "iPhone") || strings.Contains(userAgent, "iPad") || strings.Contains(userAgent, "iPod") {
		return "iOS"
	} else if strings.Contains(userAgent, "Windows") {
		return "Windows"
	} else if strings.Contains(userAgent, "Macintosh") {
		return "macOS"
	} else if strings.Contains(userAgent, "Linux") {
		return "Linux"
	} else {
		return "Unknown"
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	platform := detectPlatform(userAgent)
	fmt.Fprintf(w, "Platform detected: %s\n", platform)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
