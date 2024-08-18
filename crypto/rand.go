package crypto

import (
	"crypto/rand"
	"encoding/hex"
)

// GenRandSlice generate slice of random bytes with the specified length
func GenRandSlice(n int) ([]byte, error) {
	buf := make([]byte, n)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// GenRandStr generate random string with the specified length
func GenRandStr(n int) (randStr string, err error) {
	inputLen := n
	if n%2 == 0 {
		n = n / 2
	} else {
		n = n/2 + 1
	}

	randSl, err := GenRandSlice(n)
	if err != nil {
		return "", err
	}

	randStr = hex.EncodeToString(randSl)[:inputLen]

	return randStr, nil
}
