package cryptorand

import "crypto/rand"

const FullAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const CaptchaAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const NumbersAlphabet = "0123456789"
const CapitalLettersAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

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
