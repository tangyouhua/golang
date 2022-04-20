package main

import (
	"flag"
	"log"
)

func main() {
	ts := TinyServer{}

	debug := flag.Bool("debug", false, "enable debug info")
	flag.Parse()
	ts.debugOn(*debug)

	err := ts.ListenAndServe(":80")
	if err != nil {
		log.Fatal(err)
	}
}
