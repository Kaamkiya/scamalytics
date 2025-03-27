package main

import (
	"flag"
	"net"
	"strconv"

	"github.com/Kaamkiya/scamalytics/internal/server"
)

func main() {
	host := flag.String("host", "0.0.0.0", "The host on which to run the server")
	port := flag.Int("port", 3000, "The host on which to run the server")
	help := flag.Bool("help", false, "Show this help and exit")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	server.Run(net.JoinHostPort(*host, strconv.Itoa(*port)))
}
