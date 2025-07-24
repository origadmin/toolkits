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

	"github.com/goexts/generic"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

type Hash uint32

const (
	MD4         Hash = 1 + iota // import golang.org/x/crypto/md4
	MD5                         // import crypto/md5
	SHA1                        // import crypto/sha1
	SHA224                      // import crypto/sha256
	SHA256                      // import crypto/sha256
	SHA384                      // import crypto/sha512
	SHA512                      // import crypto/sha512
	MD5SHA1                     // no implementation; MD5+SHA1 used for TLS RSA
	RIPEMD160                   // import golang.org/x/crypto/ripemd160
	SHA3_224                    // import golang.org/x/crypto/sha3
	SHA3_256                    // import golang.org/x/crypto/sha3
	SHA3_384                    // import golang.org/x/crypto/sha3
	SHA3_512                    // import golang.org/x/crypto/sha3
	SHA512_224                  // import crypto/sha512
	SHA512_256                  // import crypto/sha512
	BLAKE2s_128                 // import golang.org/x/crypto/blake2s
	BLAKE2s_256                 // import golang.org/x/crypto/blake2s
	BLAKE2b_256                 // import golang.org/x/crypto/blake2b
	BLAKE2b_384                 // import golang.org/x/crypto/blake2b
	BLAKE2b_512                 // import golang.org/x/crypto/blake2b

	ADLER32
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
	maxHash
)

var hashEnd = maxHash
var hashes = make([]func() hash.Hash, maxHash)
var customHashNames = make(map[string]Hash)
var customNameHashes = make(map[Hash]string)

func init() {
	// Register all hash.Hash functions
	UpdateHashFunc(MD4, md4.New)
	UpdateHashFunc(MD5, md5.New)
	UpdateHashFunc(SHA1, sha1.New)
	UpdateHashFunc(SHA224, sha256.New224)
	UpdateHashFunc(SHA256, sha256.New)
	UpdateHashFunc(SHA384, sha512.New384)
	UpdateHashFunc(SHA512, sha512.New)
	UpdateHashFunc(MD5SHA1, NewMD5SHA1)
	UpdateHashFunc(RIPEMD160, ripemd160.New)
	UpdateHashFunc(SHA3_224, sha3.New224)
	UpdateHashFunc(SHA3_256, sha3.New256)
	UpdateHashFunc(SHA3_384, sha3.New384)
	UpdateHashFunc(SHA3_512, sha3.New512)
	UpdateHashFunc(SHA512_224, sha512.New512_224)
	UpdateHashFunc(SHA512_256, sha512.New512_256)
	newBlake2sHash128 := func() hash.Hash {
		return generic.Must(blake2s.New128(nil))
	}
	UpdateHashFunc(BLAKE2s_128, newBlake2sHash128)
	newBlake2sHash256 := func() hash.Hash {
		return generic.Must(blake2s.New256(nil))
	}
	UpdateHashFunc(BLAKE2s_256, newBlake2sHash256)
	newBlake2bHash256 := func() hash.Hash {
		return generic.Must(blake2b.New256(nil))
	}
	UpdateHashFunc(BLAKE2b_256, newBlake2bHash256)
	newBlake2bHash384 := func() hash.Hash {
		return generic.Must(blake2b.New384(nil))
	}
	UpdateHashFunc(BLAKE2b_384, newBlake2bHash384)
	newBlake2bHash512 := func() hash.Hash {
		return generic.Must(blake2b.New512(nil))
	}
	UpdateHashFunc(BLAKE2b_512, newBlake2bHash512)
	UpdateHashFunc(ADLER32, func() hash.Hash { return adler32.New() })
	UpdateHashFunc(CRC32, func() hash.Hash { return crc32.NewIEEE() })
	UpdateHashFunc(CRC32_ISO, func() hash.Hash { return crc32.New(crc32.MakeTable(crc32.IEEE)) })
	UpdateHashFunc(CRC32_CAST, func() hash.Hash { return crc32.New(crc32.MakeTable(crc32.Castagnoli)) })
	UpdateHashFunc(CRC32_KOOP, func() hash.Hash { return crc32.New(crc32.MakeTable(crc32.Koopman)) })
	UpdateHashFunc(CRC64_ISO, func() hash.Hash { return crc64.New(crc64.MakeTable(crc64.ISO)) })
	UpdateHashFunc(CRC64_ECMA, func() hash.Hash { return crc64.New(crc64.MakeTable(crc64.ECMA)) })
	UpdateHashFunc(FNV32, func() hash.Hash { return fnv.New32() })
	UpdateHashFunc(FNV32A, func() hash.Hash { return fnv.New32a() })
	UpdateHashFunc(FNV64, func() hash.Hash { return fnv.New64() })
	UpdateHashFunc(FNV64A, func() hash.Hash { return fnv.New64a() })
	UpdateHashFunc(FNV128, func() hash.Hash { return fnv.New128() })
	UpdateHashFunc(FNV128A, func() hash.Hash { return fnv.New128a() })
	newMapHash := func() hash.Hash {
		var mh maphash.Hash
		mh.SetSeed(maphash.MakeSeed())
		return &mh
	}
	UpdateHashFunc(MAPHASH, newMapHash)
}

func (h Hash) New() hash.Hash {
	if h > 0 && h < hashEnd {
		f := hashes[h]
		if f != nil {
			return f()
		}
	}
	panic(fmt.Sprintf("hash function %d not registered", h))
}

func (h Hash) String() string {
	if h < maxHash {
		switch h {
		case MD4:
			return "md4"
		case MD5:
			return "md5"
		case SHA1:
			return "sha1"
		case SHA224:
			return "sha224"
		case SHA256:
			return "sha256"
		case SHA384:
			return "sha384"
		case SHA512:
			return "sha512"
		case MD5SHA1:
			return "md5sha1"
		case RIPEMD160:
			return "ripemd160"
		case SHA3_224:
			return "sha3-224"
		case SHA3_256:
			return "sha3-256"
		case SHA3_384:
			return "sha3-384"
		case SHA3_512:
			return "sha3-512"
		case SHA512_224:
			return "sha512-224"
		case SHA512_256:
			return "sha512-256"
		case BLAKE2s_256:
			return "blake2s-256"
		case BLAKE2b_256:
			return "blake2b-256"
		case BLAKE2b_384:
			return "blake2b-384"
		case BLAKE2b_512:
			return "blake2b-512"
		case ADLER32:
			return "adler32"
		case CRC32:
			return "crc32"
		case CRC32_ISO:
			return "crc32-iso"
		case CRC32_CAST:
			return "crc32-cast"
		case CRC32_KOOP:
			return "crc32-koop"
		case CRC64_ISO:
			return "crc64-iso"
		case CRC64_ECMA:
			return "crc64-ecma"
		case FNV32:
			return "fnv32"
		case FNV32A:
			return "fnv32a"
		case FNV64:
			return "fnv64"
		case FNV64A:
			return "fnv64a"
		case FNV128:
			return "fnv128"
		case FNV128A:
			return "fnv128a"
		case MAPHASH:
			return "maphash"
		default:
			if name, ok := customNameHashes[h]; ok {
				return name
			}
		}
	}
	return fmt.Sprintf("Hash(%d)", h)
}

// RegisterHashFunc registers a new hash.Hash function
func RegisterHashFunc(name string, hashFunc func() hash.Hash) {
	if _, err := ParseHash(name); err == nil {
		panic(fmt.Sprintf("hash function %s already registered", name))
	}
	name = strings.ToLower(name)
	old := hashEnd
	hashEnd++
	customHashNames[name] = old
	customNameHashes[old] = name
	hashes = append(hashes, hashFunc)
}

// UpdateHashFunc updates a hash.Hash function
func UpdateHashFunc(hash Hash, hashFunc func() hash.Hash) {
	if hash >= hashEnd {
		return
	}
	hashes[hash] = hashFunc
}

// RegisterOrUpdateHashFunc registers a new hash.Hash function if it does not exist,
// otherwise updates it
func RegisterOrUpdateHashFunc(name string, hashFunc func() hash.Hash) {
	if h, err := ParseHash(name); err == nil {
		UpdateHashFunc(h, hashFunc)
	} else {
		RegisterHashFunc(name, hashFunc)
	}
}

func ParseHash(s string) (Hash, error) {
	s = strings.ToLower(s)
	if h, ok := ParseCryptoHash(s); ok {
		return h, nil
	}
	if h, ok := ParseInternalHash(s); ok {
		return h, nil
	}
	if h, ok := ParseCustomHash(s); ok {
		return h, nil
	}
	return 0, fmt.Errorf("unknown hash function: %s", s)
}

// ParseCryptoHash only deals with crypto.Hash supported algorithms
func ParseCryptoHash(s string) (Hash, bool) {
	switch s {
	case "md4":
		return MD4, true
	case "md5":
		return MD5, true
	case "sha1":
		return SHA1, true
	case "sha224":
		return SHA224, true
	case "sha256":
		return SHA256, true
	case "sha384":
		return SHA384, true
	case "sha512":
		return SHA512, true
	case "md5sha1":
		return MD5SHA1, true
	case "ripemd160":
		return RIPEMD160, true
	case "sha3-224":
		return SHA3_224, true
	case "sha3-256":
		return SHA3_256, true
	case "sha3-384":
		return SHA3_384, true
	case "sha3-512":
		return SHA3_512, true
	case "sha512-224":
		return SHA512_224, true
	case "sha512-256":
		return SHA512_256, true
	case "blake2s-256":
		return BLAKE2s_256, true
	case "blake2b-256":
		return BLAKE2b_256, true
	case "blake2b-384":
		return BLAKE2b_384, true
	case "blake2b-512":
		return BLAKE2b_512, true
	default:
		return Hash(0), false
	}
}

// ParseInternalHash algorithms that handle internal extensions
func ParseInternalHash(s string) (Hash, bool) {
	switch s {
	case "adler32":
		return ADLER32, true
	case "crc32":
		return CRC32, true
	case "crc32-iso":
		return CRC32_ISO, true
	case "crc32-cast":
		return CRC32_CAST, true
	case "crc32-koop":
		return CRC32_KOOP, true
	case "crc64-iso":
		return CRC64_ISO, true
	case "crc64-ecma":
		return CRC64_ECMA, true
	case "fnv32":
		return FNV32, true
	case "fnv32a":
		return FNV32A, true
	case "fnv64":
		return FNV64, true
	case "fnv64a":
		return FNV64A, true
	case "fnv128":
		return FNV128, true
	case "fnv128a":
		return FNV128A, true
	case "maphash":
		return MAPHASH, true
	default:
		return 0, false
	}
}

// ParseCustomHash algorithms that handle custom registrations
func ParseCustomHash(s string) (Hash, bool) {
	if h, ok := customHashNames[s]; ok {
		return h, ok
	}
	return 0, false
}

func IsCustomHash(h Hash) bool {
	return h >= maxHash && h < hashEnd
}
