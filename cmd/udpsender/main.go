package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":42069")
	if err != nil {
		log.Fatal("Failed to resolve udp addr: ", udpAddr)
	}

	conn, err := net.DialUDP(udpAddr.Network(), nil, udpAddr)
	if err != nil {
		log.Fatal("Connection failed: ", err)
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			log.Println("Error: ", err)
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Println("Failed to write line: ", err)
		}
	}
}
