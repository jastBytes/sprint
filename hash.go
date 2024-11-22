package sprint

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

// Hash returns a SHA256 hash of the given struct
func Hash[T any](s T) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(s)
	h := sha256.New()
	h.Write(b.Bytes())
	return h.Sum(nil)
}

// Hash returns a SHA256 hash of the given struct as string
func HashString[T any](s T) string {
	return fmt.Sprintf("%x\n", Hash(s))
}
