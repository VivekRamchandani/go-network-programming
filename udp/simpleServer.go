package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Create buffer for reading
	buff := make([]byte, 1024)
	meAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}

	// Create connection to listen to messages.
	conn, err := net.ListenUDP("udp4", meAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server ready.")

	// Reading received message.
	n, sender, err := conn.ReadFromUDP(buff)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s >> %s\n", sender.String(), buff[:n])
	conn.Close()
}
