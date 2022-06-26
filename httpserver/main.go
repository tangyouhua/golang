package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"
)

const VERSION = "VERSION"
const X_REAL_IP = "X-Real-IP"
const X_FORWARDED_FOR = "X-Forwarded-For"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandle)
	mux.HandleFunc("/healthz", healthzHandle)

	port := "80"
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Start HTTP server localhost:%s failed, error %s\n",
			port, err.Error())
	}
}

// index handler: 127.0.0.0/
func indexHandle(w http.ResponseWriter, r *http.Request) {
	// Copy header to response
	for key, values := range r.Header {
		for _, value := range values {
			fmt.Printf("Header key/value: [%s, %s]\n", key, value)
			w.Header().Set(key, value)
		}
	}

	// Set VERSION to header
	version := os.Getenv(VERSION)
	w.Header().Set(VERSION, version)
	fmt.Printf("os VERSION: %s\n", version)

	delayHandle(w)

	// Log client IP
	clientIP := getRemoteIP(r)
	code := http.StatusOK
	w.WriteHeader(code)
	log.Printf("Request client IP:%s, response code:%d", clientIP, code)
}

// Module10: mockup delay 0-2s
func delayHandle(w http.ResponseWriter) {
	rand.Seed(time.Now().Unix())
	delay := rand.Intn(3)
	time.Sleep(time.Duration(delay) * time.Second)
	// fmt.Fprintln(w, fmt.Sprintf("Delay %d seconds", delay))
}

// Get remote IP code
// Ref: https://juejin.cn/post/7026366595681386532
func getRemoteIP(r *http.Request) string {
	remoteAddr := r.RemoteAddr
	if ip := r.Header.Get(X_REAL_IP); ip != "" {
		remoteAddr = ip
	} else if ip = r.Header.Get(X_FORWARDED_FOR); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// health check: 127.0.0.0/healthz
func healthzHandle(w http.ResponseWriter, _ *http.Request) {
	delayHandle(w)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "HTTP server is working.")
}
