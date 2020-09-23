package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}

	lw := logWriter{}

	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
