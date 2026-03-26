package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Testing...")

	uploadPayload := make([]byte, 10000000)

	start := time.Now()

	downloadResponse, err := http.Get("https://speed.cloudflare.com/__down?bytes=10000000")
	if err != nil {
		log.Fatal(err)
	}

	bytesResult, err := io.Copy(io.Discard, downloadResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	end := time.Since(start).Seconds()

	downloadSpeed := float64(bytesResult) / end / 125000

	start = time.Now()

	_, err = http.Post("https://speed.cloudflare.com/__up", "application/octet-stream", bytes.NewReader(uploadPayload))
	if err != nil {
		log.Fatal(err)
	}

	end = time.Since(start).Seconds()

	uploadSpeed := float64(len(uploadPayload)) / end / 125000

	fmt.Printf("Download speed: %.2fMbps\n", downloadSpeed)
	fmt.Printf("Upload speed: %.2fMbps", uploadSpeed)
}
