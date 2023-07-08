package dns

import (
	"bytes"
	"testing"
)

func TestDNSHeader_Bytes(t *testing.T) {
	header := DNSHeader{
		ID:           0x1314,
		NumQuestions: 1,
	}
	expected := []byte{0x13, 0x14, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	result := header.Bytes()

	if bytes.Compare(expected, result) != 0 {
		t.Errorf("Error converting to bytes. Expected %v, got %v", expected, result)
	}
}

func TestDNSQuestion_Bytes(t *testing.T) {
	name, err := EncodeDNSName("google.com")
	if err != nil {
		t.Fatalf("Error encoding DNS name: %v", err)
	}
	question := DNSQuestion{
		Name:  name,
		Type:  TypeA,
		Class: ClassIn,
	}
	// Name field
	expected := bytes.NewBuffer([]byte{0x06})
	expected.WriteString("google")
	expected.WriteByte(0x03)
	expected.WriteString("com")
	expected.WriteByte(0x00)
	// Type and Class fields
	expected.Write([]byte{0x00, 0x01, 0x00, 0x01})

	result := question.Bytes()

	if bytes.Compare(expected.Bytes(), result) != 0 {
		t.Errorf("Error converting to bytes. Expected %v, got %v", expected.Bytes(), result)
	}
}
