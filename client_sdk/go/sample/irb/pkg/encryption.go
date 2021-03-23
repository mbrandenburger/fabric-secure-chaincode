/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

const NonceLength = 12
const KeyLength = 32

func NewEncryptionKey() []byte {
	return generateRandomBytes(KeyLength)
}

func Encrypt(plaintext []byte, key []byte) (ciphertext []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := generateRandomBytes(NonceLength)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)
	return append(nonce, ciphertext...)
}

func Decrypt(ciphertext []byte, key []byte) (plaintext []byte) {
	// split nonce and ciphertext
	nonce := ciphertext[:NonceLength]
	ciphertext = ciphertext[NonceLength:]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err = aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return plaintext
}

func generateRandomBytes(len int) []byte {
	rnd := make([]byte, len)
	if _, err := io.ReadFull(rand.Reader, rnd); err != nil {
		panic(err.Error())
	}

	return rnd
}
