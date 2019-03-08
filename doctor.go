package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	url := os.Args[1]

	if len(url) == 0 {
		log.Fatal("URL can't be blank. USAGE: doctor https://url.to.check")
	}

	var cli = &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := cli.Get(url)

	if err == nil && response.StatusCode >= 200 && response.StatusCode < 300 {
		os.Exit(0)
	}
	if err != nil {
		log.Printf("Error: %+v", err)
		os.Exit(1)
	}
	log.Printf("No Success: %+v", response)
	os.Exit(2)
}
