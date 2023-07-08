package dns

import (
	"fmt"
	"net"
	"testing"
)

func TestBuildQuery(t *testing.T) {
	query, err := BuildQuery("www.example.com", TypeA)
	if err != nil {
		t.Fatalf("Error building query: %v", err)
	}

	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		t.Fatalf("Error dialing DNS server: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write(query)
	if err != nil {
		t.Fatalf("Error querying DNS server: %v", err)
	}

	response := make([]byte, 1024)
	_, err = conn.Read(response)
	if err != nil {
		t.Fatalf("Error reading response: %v", err)
	}

	var header DNSHeader
	err = header.Parse(response[:12])
	if err != nil {
		t.Fatalf("Error parsing response header: %v", err)
	}

	fmt.Printf("Header: %+v", header)
}
