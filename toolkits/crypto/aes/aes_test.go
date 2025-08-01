/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package aes

import (
	"encoding/base64"
	"testing"
)

// SecretKey Define aes secret key 2^5
var SecretKey = []byte("2985BCFDB5FE43129843DB59825F8647")

func TestEncryptCBC(t *testing.T) {
	originalData := []byte("Hello, AES CBC!")

	ciphertext, err := EncryptCBC(originalData, SecretKey)
	if err != nil {
		t.Errorf("EncryptCBC failed: %v", err)
	}

	if len(ciphertext) == 0 {
		t.Error("EncryptCBC failed: ciphertext is empty")
	}
}

func TestEncodeCBCBase64(t *testing.T) {
	originalData := []byte("Hello, AES CBC!")

	encodedData, err := EncodeCBCBase64(originalData, SecretKey)
	if err != nil {
		t.Errorf("EncodeCBCBase64 failed: %v", err)
	}

	if len(encodedData) == 0 {
		t.Error("EncodeCBCBase64 failed: encodedData is empty")
	}

	decodedData, err := base64.RawURLEncoding.DecodeString(encodedData)
	if err != nil {
		t.Errorf("DecodeCBCBase64 failed while decoding encodedData: %v", err)
	}

	if len(decodedData) == 0 {
		t.Error("DecodeCBCBase64 failed: decodedData is empty")
	}
}

func TestDecryptCBC(t *testing.T) {
	originalData := []byte("Hello, AES CBC!")

	ciphertext, _ := EncryptCBC(originalData, SecretKey)

	decryptedData, err := DecryptCBC(ciphertext, SecretKey)
	if err != nil {
		t.Errorf("DecryptCBC failed: %v", err)
	}

	if string(decryptedData) != string(originalData) {
		t.Errorf("DecryptCBC failed: decryptedData does not match originalData")
	}
}

func TestDecodeCBCBase64(t *testing.T) {
	originalData := []byte("Hello, AES CBC!")

	encodedData, _ := EncodeCBCBase64(originalData, SecretKey)

	decryptedData, err := DecodeCBCBase64(encodedData, SecretKey)
	if err != nil {
		t.Errorf("DecodeCBCBase64 failed: %v", err)
	}

	if string(decryptedData) != string(originalData) {
		t.Errorf("DecodeCBCBase64 failed: decryptedData does not match originalData")
	}
}
