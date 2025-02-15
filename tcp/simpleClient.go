package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	serverAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
	// Connecting to Server
	conn, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	// Create a reader to read from Standard input
	reader := bufio.NewReader(os.Stdin)
	var msg string
	// Continous Sending
	for {
		msg, _ = reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent To: %s << %s\n", serverAddr.IP, msg)
		if msg == "STOP" {
			break
		}
	}
}
