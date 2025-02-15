package main

import (
	"fmt"
	"log"
	"net"
)

func connectionHandler(conn net.Conn) {
	buff := make([]byte, 1024)
	sender := conn.RemoteAddr().String()
	fmt.Println("Connected to:", sender)
	// Conitinous Reading until "STOP"
	for {
		n, err := conn.Read(buff)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Lost Connection with", sender)
				return
			}
			log.Fatal(err)
		}
		msg := string(buff[:n])
		fmt.Printf("From: %s >> %s\n", sender, msg)
		if msg == "STOP" {
			break
		}
	}

	// Close connection at the end.
	conn.Close()
}

func main() {
	// Create TCPAddr
	serverAddr, _ := net.ResolveTCPAddr("tcp", ":1234")
	// Listen to port 1234 for tcp connection request
	listener, err := net.ListenTCP("tcp", serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	// closing at the end
	defer listener.Close()
	fmt.Println("Server is Ready.")

	for {
		// Accept connection request
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go connectionHandler(conn)
	}
}
