package jwt

import "testing"

func TestGenerateKeyPair(t *testing.T) {
	pri, pub, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf(`GenerateKeyPair() = %q, %q, %v`, pri, pub, err)
	}
	t.Logf(`GenerateKeyPair() = %q, %q`, pri, pub)
}
