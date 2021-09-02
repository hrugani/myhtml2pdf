package main

import (
	"flag"
	"fmt"

	"github.com/hrugani/myhtml2pdf/webapi/app"
)

const defaultPort = 8080

func main() {
	port := flag.Int("port", defaultPort, "IP port where this server will listen")
	flag.Parse()
	fmt.Println("port reqad from input parameter = ", *port)
	app.StartApplication(*port)
}
