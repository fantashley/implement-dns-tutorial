package dns

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

type DNSHeader struct {
	ID             uint16
	Flags          uint16
	NumQuestions   uint16
	NumAnswers     uint16
	NumAuthorities uint16
	NumAdditionals uint16
}

func (h DNSHeader) Bytes() []byte {
	bytes := make([]byte, 0, 12)
	bytes = binary.BigEndian.AppendUint16(bytes, h.ID)
	bytes = binary.BigEndian.AppendUint16(bytes, h.Flags)
	bytes = binary.BigEndian.AppendUint16(bytes, h.NumQuestions)
	bytes = binary.BigEndian.AppendUint16(bytes, h.NumAnswers)
	bytes = binary.BigEndian.AppendUint16(bytes, h.NumAuthorities)
	bytes = binary.BigEndian.AppendUint16(bytes, h.NumAdditionals)
	return bytes
}

func (h *DNSHeader) Parse(data []byte) error {
	if len(data) != 12 {
		return fmt.Errorf("unexpected number of bytes: %d", len(data))
	}
	var header DNSHeader
	var i int
	header.ID = binary.BigEndian.Uint16(data[i : i+2])
	i += 2
	header.Flags = binary.BigEndian.Uint16(data[i : i+2])
	i += 2
	header.NumQuestions = binary.BigEndian.Uint16(data[i : i+2])
	i += 2
	header.NumAnswers = binary.BigEndian.Uint16(data[i : i+2])
	i += 2
	header.NumAuthorities = binary.BigEndian.Uint16(data[i : i+2])
	i += 2
	header.NumAdditionals = binary.BigEndian.Uint16(data[i : i+2])

	*h = header
	return nil
}

func EncodeDNSName(name string) ([]byte, error) {
	buf := &bytes.Buffer{}
	parts := strings.FieldsFunc(name, func(r rune) bool { return r == '.' })
	for _, part := range parts {
		length := uint8(len([]byte(part)))
		err := binary.Write(buf, binary.BigEndian, length)
		if err != nil {
			return nil, fmt.Errorf("error writing length of part %q to buffer: %w", part, err)
		}
		_, err = buf.WriteString(part)
		if err != nil {
			return nil, fmt.Errorf("error writing part %q to buffer: %w", part, err)
		}
	}
	err := buf.WriteByte(0x00)
	if err != nil {
		return nil, fmt.Errorf("error writing null to end of buffer: %w", err)
	}
	return buf.Bytes(), nil
}

type DNSQuestion struct {
	Name  []byte
	Type  uint16
	Class uint16
}

func (q DNSQuestion) Bytes() []byte {
	bytes := make([]byte, 0, len(q.Name)+4)
	bytes = append(bytes, q.Name...)
	bytes = binary.BigEndian.AppendUint16(bytes, q.Type)
	bytes = binary.BigEndian.AppendUint16(bytes, q.Class)
	return bytes
}

type DNSRecord struct {
	Name  []byte
	Type  uint16
	Class uint16
	TTL   uint32
	Data  []byte
}
