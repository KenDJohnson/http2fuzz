package main

import (
	"flag"
	"os"

	"github.com/johnso51/http2fuzz/config"
	"github.com/johnso51/http2fuzz/fuzzer"
)

func main() {

	if config.FuzzMode == config.ModeClient {
		fuzzer.Client()
	} else if config.FuzzMode == config.ModeServer {
		fuzzer.Server()
	} else {
		flag.Usage()
		os.Exit(1)
	}

	select {}
}
