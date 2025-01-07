package cryptorand

import "crypto/rand"

// FullAlphabet defines one of possible alphabets used for generating random string
const FullAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// CaptchaAlphabet defines one of possible alphabets used for generating random string
const CaptchaAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// NumbersAlphabet defines one of possible alphabets used for generating random string
const NumbersAlphabet = "0123456789"

// CapitalLettersAlphabet defines one of possible alphabets used for generating random string
const CapitalLettersAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandomBytes generates cryptographically secure random bytes
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandomString generates cryptographically secure random string using alphabet provided
func GenerateRandomString(alphabet string, n int) (string, error) {
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = alphabet[b%byte(len(alphabet))]
	}
	return string(bytes), nil
}
