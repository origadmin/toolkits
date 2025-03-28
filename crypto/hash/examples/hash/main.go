/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package main

import (
	"fmt"
	"log"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func main() {
	// 创建配置
	config := &types.Config{
		Algorithm:  types.TypeArgon2,
		TimeCost:   3,
		MemoryCost: 64 * 1024,
		Threads:    4,
		SaltLength: 16,
	}

	// 创建加密实例
	crypto, err := core.NewCrypto(config)
	if err != nil {
		log.Fatal(err)
	}

	// 测试密码
	password := "test123"

	// 生成哈希
	hashed, err := crypto.Hash(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Generated hash: %s\n", hashed)

	// 验证密码
	err = crypto.Verify(hashed, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Password verified successfully!")

	// 测试错误的密码
	wrongPassword := "wrong123"
	err = crypto.Verify(hashed, wrongPassword)
	if err != nil {
		fmt.Println("Wrong password detected as expected")
	}
}
