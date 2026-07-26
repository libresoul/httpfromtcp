package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("Failed to create tcp socket: ", err)
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Failed to accept connection to the listener: ", err)
		}
		fmt.Println("Connection accepted")

		for line := range getLinesChannel(conn) {
			fmt.Println(line)
		}
		fmt.Println("Connection has been closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		defer f.Close()

		var currentLine string

		for {
			buffer := make([]byte, 8)
			_, err := f.Read(buffer)

			if !strings.Contains(string(buffer), "\n") {
				currentLine += string(buffer)
			}
			parts := strings.Split(string(buffer), "\n")

			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				log.Fatal(err)
			}

			for j, p := range parts {
				if len(parts) == 1 {
					continue
				}
				if j != len(parts)-1 {
					ch <- currentLine + p
					currentLine = ""
					continue
				}
				currentLine += p
			}
		}
	}()

	return ch
}
