package main

import (
	"flag"
	"log"
	"os"
	"strconv"
)

func main() {
	ts := TinyServer{}

	debug := flag.Bool("debug", false, "enable debug info")
	flag.Parse()
	ts.debugOn(*debug)

	err := ts.ListenAndServe(port())
	if err != nil {
		log.Fatal(err)
	}
}

func port() string {
	portstr := os.Getenv("TYH_SERVICE_PORT")
	if portstr != "" {
		_, err := strconv.ParseInt(portstr, 10, 64)
		if err == nil {
			return ":" + portstr
		}
	}
	return ":80"
}
