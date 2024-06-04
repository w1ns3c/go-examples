package crypto

import (
	"crypto/rand"
	"encoding/hex"
)

func GenRandSlice(n int) (buf []byte, err error) {
	buf = make([]byte, n)
	_, err = rand.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func GenRandStr(n int) (s string, err error) {
	n_init := n
	if n%2 == 0 {
		n = n / 2
	} else {
		n = n/2 + 1
	}
	buf, err := GenRandSlice(n)
	if err != nil {
		return "", err
	}

	s = hex.EncodeToString(buf)
	return s[:n_init], nil
}
