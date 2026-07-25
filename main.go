package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Failed to open file: ", err)
	}

	defer file.Close()

	for {
		buffer := make([]byte, 8)
		_, err := file.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("read: %s\n", string(buffer))
	}
}
