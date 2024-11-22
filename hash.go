package sprint

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

// Hash returns a SHA256 hash of the given struct
func Hash[T any](s T) ([]byte, error) {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(s); err != nil {
		return nil, err
	}
	h := sha256.New()
	h.Write(b.Bytes())
	return h.Sum(nil), nil
}

// Hash returns a SHA256 hash of the given struct as string
func HashString[T any](s T) (string, error) {
	h, err := Hash(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x\n", h), nil
}
