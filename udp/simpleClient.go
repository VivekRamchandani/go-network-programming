package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	// Creating a UDP port to send message.
	conn, err := net.DialUDP("udp4", nil, serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	// Sending message to server
	_, err = conn.Write([]byte("Hello Server"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s << Hello Word\n", serverAddr.String())
	// Closing connection
	conn.Close()
}
