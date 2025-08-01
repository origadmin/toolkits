/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package aes provides the AES encryption and decryption
package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
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
// Deprecated: This function uses a static IV derived from the key, which is insecure.
// It is provided for compatibility with legacy systems (e.g., WeChat Pay) and should not be used for new encryption implementations.
// For secure encryption, use a function that supports random IVs, such as AES-GCM.
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
// Deprecated: This function uses a static IV derived from the key, which is insecure.
// It is provided for compatibility with legacy systems (e.g., WeChat Pay) and should not be used for new encryption implementations.
// For secure decryption, use a function that supports random IVs, such as AES-GCM.
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

// EncryptGCM encrypts data using AES-GCM, a modern and secure authenticated encryption mode.
// It generates a random nonce, prepends it to the ciphertext, and returns the result.
// The key must be 16, 24, or 32 bytes long to select AES-128, AES-192, or AES-256.
func EncryptGCM(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Seal will append the output to the first argument; we pass nonce as the first argument
	// to prepend it to the ciphertext.
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// DecryptGCM decrypts data that was encrypted using AES-GCM.
// It expects the nonce to be prepended to the ciphertext.
func DecryptGCM(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
