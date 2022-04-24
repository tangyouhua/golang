package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

type TinyServer struct {
	debug bool
}

func (ts *TinyServer) ListenAndServe(addr string) error {
	http.HandleFunc("/healthz", ts.healthzHandler)
	http.HandleFunc("/", ts.defaultHandler)
	return http.ListenAndServe(addr, nil)
}

func (ts *TinyServer) healthzHandler(w http.ResponseWriter, r *http.Request) {
	copyHeader(w, r)
	setVersion(w, r)
	statusCode := http.StatusOK
	ts.Println("Client IP:", r.RemoteAddr, ", HTTP code:", statusCode)
	w.WriteHeader(statusCode)
}

func (ts *TinyServer) defaultHandler(w http.ResponseWriter, r *http.Request) {
	copyHeader(w, r)
	setVersion(w, r)
	statusCode := http.StatusForbidden
	ts.Println("Client IP:", r.RemoteAddr, ", HTTP code:", statusCode)
	w.WriteHeader(statusCode)
}

func copyHeader(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		w.Header().Add(key, strings.Join(value, ","))
	}
}

func setVersion(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	w.Header().Add("VERSION", version)
}

func (ts *TinyServer) Println(v ...interface{}) {
	if ts.debug {
		log.Println(v)
	}
}

func (ts *TinyServer) debugOn(value bool) {
	ts.debug = value
}
