package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"net/url"
)

func main() {
	server := flag.String("server", "localhost:8080", "the server to use")

	address := url.URL{Scheme: "http", Host: *server, Path: "/certificate"}

	resp, err := http.Get(address.String())
	if err != nil {
		log.Fatal(err)
	}

	certificateResponse := CertificateResponse{}
	json.NewDecoder(resp.Body).Decode(&certificateResponse)

	log.Printf("Response: %+v", certificateResponse)
}
