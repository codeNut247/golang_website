package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Server stores Hostname and Http Port number
type Server struct {
	Hostname string `json:"Hostname"`
	HTTPPort int    `json:"HTTPPort"`
}

// Starts http listener (later could start https as well)
func run(handlers http.Handler, s Server) {
	// TODO: httpStart(handler, server)
	if s.HTTPPort != 0 {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTP "+httpAddress(s))
		// Start Http Listener
		log.Fatal(http.ListenAndServe(httpAddress(s), handlers))
	} else {
		log.Println("Config File doesn't specify a Portnumber")
	}
}

// httpAddress returns httpAddress
func httpAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPPort)
}
