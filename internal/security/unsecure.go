package security

//nolint:gosec // It's hard to change now without breaking. Issue #44.
import (
	"crypto/sha1"
	"fmt"
)

// UnsecureSHA1 is not secure, but makes a fixed length password.
// With v2, I'm planning to move away from it, but that would break
// all existing vaults, so I have to be careful and make sure a proper
// migration script/function exists.
func UnsecureSHA1(text string) []byte {
	result := make([]byte, passwordHashLength)

	hash := sha1.New() //nolint:gosec // yolo?
	_, _ = hash.Write([]byte(text))
	h := hash.Sum(nil)
	text = fmt.Sprintf("%x", h)

	copy(result, text[0:passwordHashLength])

	return result
}
