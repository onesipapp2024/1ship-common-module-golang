package helper

import (
	"testing"
)

func TestGenerateRandomStringWithLength(t *testing.T) {
	length := 10
	random, err := GenerateRandomStringWithLength(length)
	if err != nil {
		t.Fatalf(`GenerateRandomStringWithLength(%v) = %q, %v`, length, random, err)
	}
}

func TestParseClientType(t *testing.T) {
	ua := "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1"
	clientType := ParseClientType(ua)
	if clientType != ClientTypeIOS {
		t.Fatalf(`ParseClientType(%q) = %v`, ua, clientType)
	}
}
