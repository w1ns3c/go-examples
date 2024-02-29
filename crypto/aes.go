package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

const (
	AES_KEY_LEN = 32
)

// EncryptAES is example func for encrypt data with AES cipher
// key len must be 16 | 24 | 32 signs
// Note:
//
//	use FillKeyWithZero or FillKeyWithHash or FillKeyWithRand
//	for extend key to AES_KEY_LEN
func EncryptAES(data []byte, key []byte) (cData []byte, err error) {
	keyBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("can't create new cipher from KEY: %v", err)
	}

	gcm, err := cipher.NewGCM(keyBlock)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	cData = gcm.Seal(nonce, nonce, data, nil)

	return cData, err
}

// DecryptAES is example func for decrypt data with AES cipher
// key len must be 16 | 24 | 32 signs
// Note:
//
//	use FillKeyWithZero or FillKeyWithHash or FillKeyWithRand
//	for extend key to AES_KEY_LEN
func DecryptAES(cData []byte, key []byte) (data []byte, err error) {
	keyBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("can't create new cipher from KEY: %v", err)
	}

	gcm, err := cipher.NewGCM(keyBlock)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, cData := cData[:nonceSize], cData[nonceSize:]

	return gcm.Open(nil, nonce, cData, nil)
}

func FillKeyWithZero(key []byte) []byte {
	if len(key) == AES_KEY_LEN {
		return key
	}

	filledKey := make([]byte, AES_KEY_LEN)
	for i := 0; i < len(key); i++ {
		filledKey[i] = key[i]
	}

	for i := len(key); i < len(filledKey); i++ {
		filledKey[i] = 0
	}

	return filledKey
}

func FillKeyWithHash(key []byte) []byte {
	hash := sha256.Sum256(key)
	return hash[:]
}

func FillKeyWithRand(key []byte) (filledKey []byte, err error) {
	if len(key) == AES_KEY_LEN {
		return key, nil
	}

	filledKey = make([]byte, AES_KEY_LEN)
	for i := 0; i < len(key); i++ {
		filledKey[i] = key[i]
	}

	buf, err := GenRandSlice(AES_KEY_LEN - len(key))
	if err != nil {
		return nil, err
	}
	if len(key)+len(buf) != len(filledKey) {
		return nil, fmt.Errorf("wrong len of generated buffer")
	}

	for i, j := len(key), 0; i < len(filledKey); i, j = i+1, j+1 {
		filledKey[i] = buf[j]
	}

	return filledKey, nil
}
