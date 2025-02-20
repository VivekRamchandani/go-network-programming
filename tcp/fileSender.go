package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func connectionHandler(conn *net.TCPConn) {
	buff := make([]byte, 128)
	file, err := os.Open("./_files/send.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer conn.Close()
	for {
		n, err := file.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		conn.Write(buff[:n])
	}
}

func main() {
	me, err := net.ResolveTCPAddr("tcp4", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp4", me)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("Server is ready.")

	for {
		conn, err := listener.AcceptTCP()
		fmt.Println("Connected:", conn.RemoteAddr().String())
		if err != nil {
			log.Fatal(err)
		}
		go connectionHandler(conn)
	}
}
