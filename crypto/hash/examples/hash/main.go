/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package main

import (
	"fmt"
	"log"

	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/argon2"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func main() {

	// reate cryptographic instance
	crypto, err := hash.NewCrypto(types.TypeArgon2, func(config *types.Config) {
		config.SaltLength = 16
		config.ParamConfig = argon2.DefaultParams().String()
	})
	if err != nil {
		log.Fatal(err)
	}

	// Test password
	password := "test123"

	// Generate hash
	hashed, err := crypto.Hash(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Generated hash: %s\n", hashed)

	//  Verify password
	err = crypto.Verify(hashed, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Password verified successfully!")

	// Test wrong password
	wrongPassword := "wrong123"
	err = crypto.Verify(hashed, wrongPassword)
	if err != nil {
		fmt.Println("Wrong password detected as expected")
	}
}
