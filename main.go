package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Certificate represents the structure of our JSON response
type Certificate struct {
	IssuerCAID     int64  `json:"issuer_ca_id"`
	IssuerName     string `json:"issuer_name"`
	CommonName     string `json:"common_name"`
	NameValue      string `json:"name_value"`
	ID             int64  `json:"id"`
	EntryTimestamp string `json:"entry_timestamp"`
	NotBefore      string `json:"not_before"`
	NotAfter       string `json:"not_after"`
	SerialNumber   string `json:"serial_number"`
	ResultCount    int    `json:"result_count"`
}

func main() {
	// Parsing domain from command line arguments
	domain := flag.String("domain", "", "Domain to search for SSL/TLS certificates")
	flag.Parse()

	if *domain == "" {
		log.Fatal("Domain is required. Usage: go run main.go -domain=insecure.com")
	}

	// Calling crt.sh API
	certificates, err := queryCertificates(*domain)
	if err != nil {
		log.Fatalf("Failed to query crt.sh: %v", err)
	}

	// Marshaling entire slice of certificates to JSON
	certsJSON, err := json.MarshalIndent(certificates, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal certificates to JSON: %v", err)
	}

	// Printing out the JSON string of certificates
	fmt.Println(string(certsJSON))
}

func queryCertificates(domain string) ([]Certificate, error) {
	// crt.sh URL. Note: This uses an undocumented API endpoint that might change.
	apiURL := "https://crt.sh/?q=%25." + url.QueryEscape(domain) + "&output=json"

	// Sending HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Reading response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parsing JSON response
	var certificates []Certificate
	if err := json.Unmarshal(body, &certificates); err != nil {
		return nil, err
	}

	return certificates, nil
}
