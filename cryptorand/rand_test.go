package cryptorand

import "testing"

func TestGenerateRandomString(t *testing.T) {
	a, err := GenerateRandomString("a", 5)
	if err != nil {
		t.Error(err)
		return
	}
	if a != "aaaaa" {
		t.Error("not works")
	}
}
