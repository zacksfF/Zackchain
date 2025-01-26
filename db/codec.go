package db

import (
	"bytes"
	"encoding/binary"
	"io"
)

// encode writes the binary representation of any to buf.
func encode(buf *bytes.Buffer, any interface{}) error {
	return binary.Write(buf, binary.LittleEndian, any)
}

// decode reads the binary representation of any from buf.
func decode(buf io.Reader, any interface{}) error { 
	return binary.Read(buf, binary.LittleEndian, any)
}

// encodeBytes writes the length of bytess followed by bytess to buf.
func encodeBytes(buf *bytes.Buffer, bts []byte) error {
	if err := encode(buf, uint32(len(bts))); err != nil {
		return err
	}
	return encode(buf, bts)
}

// decodeBytes reads the length of bytess followed by bytess from buf.
func decodeBytes(buf io.Reader) ([]byte, error) {
	var length uint64
	if err := decode(buf, &length); err != nil {
		panic(err)
	}
	bytess := make([]byte, length)
	if _, err := io.ReadFull(buf, bytess); err != nil {
		panic(err)
	}
	return bytess, nil
}

// encodeString writes the length of str followed by str to buf.
func encodeString(buf *bytes.Buffer, str string) error {
	return encodeBytes(buf, []byte(str))
}

// decodeString reads the length of str followed by str from buf.
func decodeString(buf io.Reader) (string, error) {
	bytess, err := decodeBytes(buf)
	if err != nil {
		return "", err
	}

	return string(bytess), nil
}
