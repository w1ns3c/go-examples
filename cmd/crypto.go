package main

import (
	"fmt"
	"strings"

	"github.com/w1ns3c/go-examples/crypto"
)

const (
	AES_KEY = "my_aes_key"
)

func main() {
	data := []byte(strings.Repeat(`some supersecret data`, 4))
	//data := []byte(`some supersecret data`)

	fmt.Println("--- Example RSA ---")
	ExampleRSA(data)
	fmt.Print("\n\n")

	fmt.Println("--- Example AES ---")
	key := []byte(AES_KEY)
	ExampleAES(data, key)
}

func ExampleAES(data []byte, key []byte) {
	fmt.Println(key)
	key, err := crypto.FillKeyWithRand(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(key)

	cDataAES, err := crypto.EncryptAES(data, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cDataAES)
	fmt.Println(string(cDataAES))

	dataAES, err := crypto.DecryptAES(cDataAES, key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(dataAES)

	fmt.Println("Original:  ", string(data))
	fmt.Println("Decrypted: ", string(dataAES))
}

func ExampleRSA(data []byte) {
	priv, pub, err := crypto.GenKeys(2048)
	if err != nil {
		fmt.Println("Can't generate priv/pub key pair")
		fmt.Println(err)
		return
	}

	//fmt.Println(string(priv), "\n", string(pub))

	privKey, err := crypto.ReadPrivateKey(priv)
	if err != nil {
		fmt.Printf("Can't parse priv key: %v\n", err)
		return
	}

	pubKey, err := crypto.ReadPubKey(pub)
	if err != nil {
		fmt.Printf("Can't parse pub key: %v\n", err)
		return
	}

	cData, err := crypto.EncryptRSA(data, pubKey)
	if err != nil {
		fmt.Printf("Can't encrypt data: %v\n", err)
		return
	}

	resData, err := crypto.DecryptRSA(cData, privKey)
	if err != nil {
		fmt.Printf("Can't decrypt data: %v\n", err)
		return
	}

	fmt.Println("Original:  ", string(data))
	fmt.Println("Decrypted: ", string(resData))
}
