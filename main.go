package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

type OSResponse struct {
    Platform  string `json:"platform"`
    OS        string `json:"os"`
    UserAgent string `json:"user_agent"`
    Device    string `json:"device"`
}

func detectOS(userAgent string) string {
    if strings.Contains(userAgent, "Windows") {
        return "Windows"
    } else if strings.Contains(userAgent, "Macintosh") {
        return "macOS"
    } else if strings.Contains(userAgent, "Linux") {
        return "Linux"
    } else if strings.Contains(userAgent, "iPhone") || strings.Contains(userAgent, "iPad") {
        return "iOS"
    } else if strings.Contains(userAgent, "Android") {
        return "Android"
    } else if strings.Contains(userAgent, "Chrome OS") {
        return "Chrome OS"
    }
    return "Unknown"
}

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
    }
    return "Unknown"
}

func detectDevice(userAgent string) string {
    if strings.Contains(userAgent, "Mobile") {
        return "Mobile"
    } else if strings.Contains(userAgent, "Tablet") {
        return "Tablet"
    } else {
        return "Desktop"
    }
}

func platformAPIHandler(w http.ResponseWriter, r *http.Request) {
    userAgent := r.Header.Get("User-Agent")
    platform := detectPlatform(userAgent)
    os := detectOS(userAgent)
    device := detectDevice(userAgent)

    response := OSResponse{Platform: platform, OS: os, UserAgent: userAgent, Device: device}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func main() {
    http.HandleFunc("/", homePageHandler)
    http.HandleFunc("/detect-platform", platformAPIHandler)

    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
