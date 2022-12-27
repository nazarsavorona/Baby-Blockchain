package Hash

import (
	"crypto/sha256"
	"fmt"
)

func ToSHA256(message string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(message)))
}
