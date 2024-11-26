/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package aes provides the AES encryption and decryption
package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// PKCS5Padding adds padding to the plaintext based on the blockSize.
//
// plaintext: The data to pad.
// blockSize: The block size used for padding.
// []byte: The padded plaintext.
func PKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

// PKCS5UnPadding removes the padding from the given byte array.
//
// Parameters:
// - data: the byte array to remove the padding from.
//
// Returns:
// - []byte: the byte array with the padding removed.
func PKCS5UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// EncryptCBC encrypts the given original data using the provided key using the AES encryption algorithm in CBC mode.
//
// Parameters:
// - data: the data to be encrypted ([]byte)
// - key: the encryption key ([]byte)
//
// Returns:
// - []byte: the encrypted data
// - error: an error if the encryption fails
func EncryptCBC(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	data = PKCS5Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	return crypted, nil
}

// EncodeCBCBase64 encrypts the given original data using the provided key and returns the encrypted data as a base64-encoded string.
//
// Parameters:
// - data: the data to be encrypted ([]byte)
// - key: the encryption key ([]byte)
//
// Returns:
// - string: the encrypted data as a base64-encoded string
// - error: an error if the encryption fails
func EncodeCBCBase64(data, key []byte) (string, error) {
	crypted, err := EncryptCBC(data, key)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(crypted), nil
}

// DecryptCBC decrypts the given ciphertext using the provided key.
//
// Parameters:
// - crypted: the ciphertext to be decrypted ([]byte)
// - key: the decryption key ([]byte)
//
// Returns:
// - []byte: the decrypted data
// - error: an error if the decryption fails
func DecryptCBC(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	data := make([]byte, len(crypted))
	blockMode.CryptBlocks(data, crypted)
	data = PKCS5UnPadding(data)
	return data, nil
}

// DecodeCBCBase64 decrypts the base64-encoded data using the provided key and returns the decrypted data.
//
// Parameters:
// - data: the base64-encoded data to be decrypted (string)
// - key: the decryption key ([]byte)
//
// Returns:
// - []byte: the decrypted data
// - error: an error if the decryption fails
func DecodeCBCBase64(data string, key []byte) ([]byte, error) {
	crypted, err := base64.RawURLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return DecryptCBC(crypted, key)
}
