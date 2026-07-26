package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Failed to open file: ", err)
	}

	lineCh := getLinesChannel(file)
	for line := range lineCh {
		fmt.Println("read:", line)
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
