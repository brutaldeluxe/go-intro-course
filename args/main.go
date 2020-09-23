package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var args = os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Missing arguments")
		os.Exit(1)
	}

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
