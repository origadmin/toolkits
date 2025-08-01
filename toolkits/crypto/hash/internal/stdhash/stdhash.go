/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package stdhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"hash/maphash"
	"strings"

	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

type Hash uint32

// crypto hashes
const (
	MD4        Hash = 1 + iota // import golang.org/x/crypto/md4
	MD5                        // import crypto/md5
	SHA1                       // import crypto/sha1
	SHA224                     // import crypto/sha256
	SHA256                     // import crypto/sha256
	SHA384                     // import crypto/sha512
	SHA512                     // import crypto/sha512
	MD5SHA1                    // no implementation; MD5+SHA1 used for TLS RSA
	RIPEMD160                  // import golang.org/x/crypto/ripemd160
	SHA3_224                   // import golang.org/x/crypto/sha3
	SHA3_256                   // import golang.org/x/crypto/sha3
	SHA3_384                   // import golang.org/x/crypto/sha3
	SHA3_512                   // import golang.org/x/crypto/sha3
	SHA512_224                 // import crypto/sha512
	SHA512_256                 // import crypto/sha512
	cryptoHashEnd
)

// internal hashes
const (
	ADLER32 = cryptoHashEnd + iota
	CRC32
	CRC32_ISO
	CRC32_CAST
	CRC32_KOOP
	CRC64_ISO
	CRC64_ECMA
	FNV32
	FNV32A
	FNV64
	FNV64A
	FNV128
	FNV128A
	MAPHASH
	internalHashEnd
)

// end of crypto hashes
const (
	maxHash           = internalHashEnd
	firstCryptoHash   = MD4
	firstInternalHash = ADLER32
)

var hashes = make([]func() hash.Hash, maxHash)
var hashNames = make([]string, maxHash)
var hashNameMap = make(map[string]Hash)

func init() {
	// Register all hash.Hash functions
	registerHash(MD4, "md4", md4.New)
	registerHash(MD5, "md5", md5.New)
	registerHash(SHA1, "sha1", sha1.New)
	registerHash(SHA224, "sha224", sha256.New224)
	registerHash(SHA256, "sha256", sha256.New)
	registerHash(SHA384, "sha384", sha512.New384)
	registerHash(SHA512, "sha512", sha512.New)
	registerHash(MD5SHA1, "md5+sha1", NewMD5SHA1)
	registerHash(RIPEMD160, "ripemd-160", ripemd160.New)
	registerHash(SHA3_224, "sha3-224", sha3.New224)
	registerHash(SHA3_256, "sha3-256", sha3.New256)
	registerHash(SHA3_384, "sha3-384", sha3.New384)
	registerHash(SHA3_512, "sha3-512", sha3.New512)
	registerHash(SHA512_224, "sha512/224", sha512.New512_224)
	registerHash(SHA512_256, "sha512/256", sha512.New512_256)
	// BLAKE2 hashes are handled by the blake2 package directly
	// and are not registered here with a New() function.
	// Their Hash enum values are still defined for classification.
	registerHash(ADLER32, "adler32", func() hash.Hash { return adler32.New() })
	registerHash(CRC32, "crc32", func() hash.Hash { return crc32.NewIEEE() })
	registerHash(CRC32_ISO, "crc32-iso", func() hash.Hash { return crc32.New(crc32.MakeTable(crc32.IEEE)) })
	registerHash(CRC32_CAST, "crc32-cast", func() hash.Hash { return crc32.New(crc32.MakeTable(crc32.Castagnoli)) })
	registerHash(CRC32_KOOP, "crc32-koop", func() hash.Hash { return crc32.New(crc32.MakeTable(crc32.Koopman)) })
	registerHash(CRC64_ISO, "crc64-iso", func() hash.Hash { return crc64.New(crc64.MakeTable(crc64.ISO)) })
	registerHash(CRC64_ECMA, "crc64-ecma", func() hash.Hash { return crc64.New(crc64.MakeTable(crc64.ECMA)) })
	registerHash(FNV32, "fnv32", func() hash.Hash { return fnv.New32() })
	registerHash(FNV32A, "fnv32a", func() hash.Hash { return fnv.New32a() })
	registerHash(FNV64, "fnv64", func() hash.Hash { return fnv.New64() })
	registerHash(FNV64A, "fnv64a", func() hash.Hash { return fnv.New64a() })
	registerHash(FNV128, "fnv128", func() hash.Hash { return fnv.New128() })
	registerHash(FNV128A, "fnv128a", func() hash.Hash { return fnv.New128a() })
	newMapHash := func() hash.Hash {
		var mh maphash.Hash
		mh.SetSeed(maphash.MakeSeed())
		return &mh
	}
	registerHash(MAPHASH, "maphash", newMapHash)
}

func (h Hash) New() hash.Hash {
	if h > 0 && h < maxHash {
		f := hashes[h]
		if f != nil {
			return f()
		}
	}
	panic(fmt.Sprintf("hash function %d not registered", h))
}

func (h Hash) String() string {
	if h > 0 && h < maxHash {
		return hashNames[h]
	}
	return fmt.Sprintf("Hash(%d)", h)
}

func (h Hash) IsCrypto() bool {
	return h >= firstCryptoHash && h < cryptoHashEnd
}

func (h Hash) IsInternal() bool {
	return h >= firstInternalHash && h < internalHashEnd
}

// UpdateHashFunc updates a hash.Hash function
func UpdateHashFunc(hash Hash, hashFunc func() hash.Hash) {
	if hash >= maxHash {
		return
	}
	hashes[hash] = hashFunc
}

// ParseHash parses a hash function name and returns the Hash type
func ParseHash(s string) (Hash, error) {
	s = strings.ToLower(s)
	if h, ok := hashNameMap[s]; ok {
		return h, nil
	}

	return 0, fmt.Errorf("unknown hash function: %s", s)
}

func registerHash(h Hash, name string, newFunc func() hash.Hash) {
	UpdateHashFunc(h, newFunc)
	hashNames[h] = name
	hashNameMap[name] = h
}
