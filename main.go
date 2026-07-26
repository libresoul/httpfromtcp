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

	defer file.Close()
	var currentLine string

	for {
		buffer := make([]byte, 8)
		_, err := file.Read(buffer)

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
				fmt.Printf("read: %s\n", currentLine+p)
				currentLine = ""
				continue
			}
			currentLine += p
		}

	}
}
