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
    fmt.Print("Enter Remote Addr: ")
    reader := bufio.NewReader(os.Stdin)
    addrStr, err := reader.ReadString('\n')
    addrStr = strings.TrimRight(addrStr, "\n")
	remote, err := net.ResolveTCPAddr("tcp", addrStr + ":9999")
    if err != nil { 
        log.Fatal(err)
    }

    conn, err := net.DialTCP("tcp", nil, remote)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    buff := make([]byte, 128)
    file, err := os.OpenFile("./_files/recieve.txt", os.O_WRONLY | os.O_CREATE, 0755)
    fmt.Println("File Created.")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    for {
        n, err := conn.Read(buff)
        if err != nil {
            if err.Error() == "EOF" {
                fmt.Println("Disconnected")
                break
            }
            log.Fatal(err)
        }
        _, err = file.Write(buff[:n])
        if err != nil {
            log.Fatal(err)
        }
    }
}