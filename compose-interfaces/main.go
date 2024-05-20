package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	fmt.Println("Compose Interfaces!")
	payload := []byte("Clint")
	hashAndBroadCast(NewHashReader(payload))
}

// Interface defined to compose io.Reader as a type of hashReader
type HashReader interface {
	io.Reader
	hash() string
}

type hashReader struct {
	bytes.Reader               // Used for the IO Read
	buff         *bytes.Buffer // Used to read byte data and create hash
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: *bytes.NewReader(b),
		buff:   bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	hash := sha1.Sum(h.buff.Bytes())
	return hex.EncodeToString(hash[:])
}

func hashAndBroadCast(r HashReader) error {
	// hash := r.(*hashReader).hash()
	hash := r.hash()

	fmt.Println(hash)
	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("String of bytes: ", string(b))
	return nil
}
