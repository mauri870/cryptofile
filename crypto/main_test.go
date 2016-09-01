package crypto

import (
	"testing"
)

var (
	stubs = []string{"Hello World", "The Quick Brown Fox", "The difficult we do immediately; the impossible takes a little longer."}
)

func TestEncryptDecrypt(t *testing.T) {
	key := []byte("hDmPPK2b76ROBqd4uWQcu0ruUyz0tVXd")
	for _, text := range stubs {
		ciphertext, err := Encrypt(key, []byte(text))
		if err != nil {
			t.Error(err)
		}

		decryptedText, err := Decrypt(key, ciphertext)
		if err != nil {
			t.Error(err)
		}

		if string(decryptedText) != text {
			t.Errorf("Expect %s but got %s on decryption", text, decryptedText)
		}
	}
}
