package dns

import (
	"fmt"
	"math/rand"
)

const TypeA uint16 = 1
const ClassIn uint16 = 1

func BuildQuery(domainName string, recordType uint16) ([]byte, error) {
	name, err := EncodeDNSName(domainName)
	if err != nil {
		return nil, fmt.Errorf("error encoding domain name %q: %w", domainName, err)
	}

	id := uint16(rand.Uint32())

	recursionDesired := uint16(1 << 8)

	header := DNSHeader{
		ID:           id,
		Flags:        recursionDesired,
		NumQuestions: 1,
	}
	question := DNSQuestion{
		Name:  name,
		Type:  recordType,
		Class: ClassIn,
	}
	return append(header.Bytes(), question.Bytes()...), nil
}
