package core

import (
	"fmt"
	"hash"
	"testing"
)

type CustomHash struct {
}

func (c *CustomHash) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (c *CustomHash) Reset() {
}

func (c *CustomHash) Size() int {
	return 0
}

func (c *CustomHash) BlockSize() int {
	return 0
}

func (c *CustomHash) Sum(b []byte) []byte {
	return []byte("mock_custom_hash")
}

type CustomHash2 struct {
}

func (c *CustomHash2) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (c *CustomHash2) Reset() {

}

func (c *CustomHash2) Size() int {
	return 0
}

func (c *CustomHash2) BlockSize() int {
	return 0
}

func (c *CustomHash2) Sum(b []byte) []byte {
	return []byte("custom_hash_2")
}

type CustomHash3 struct {
	counter int
	sum     []byte
}

func (c *CustomHash3) Write(p []byte) (n int, err error) {
	c.counter += len(p)
	return len(p), nil
}

func (c *CustomHash3) Sum(b []byte) []byte {
	return c.sum
}

func (c *CustomHash3) Reset() {
	c.counter = 0
}

func (c *CustomHash3) Size() int {
	return 16
}

func (c *CustomHash3) BlockSize() int {
	return 64
}

// 注册自定义哈希算法
func init() {
	RegisterHashFunc("CUSTOM", func() hash.Hash {
		return &CustomHash{}
	})

	RegisterHashFunc("CUSTOM2", func() hash.Hash {
		return &CustomHash2{}
	})

	RegisterOrUpdateHashFunc("CONFLICT", func() hash.Hash {
		return &CustomHash3{sum: []byte("version1")}
	})

	RegisterOrUpdateHashFunc("CONFLICT", func() hash.Hash {
		return &CustomHash3{sum: []byte("version2")}
	})
}

func TestHashAlgorithms(t *testing.T) {
	tests := []struct {
		name        string
		hash        Hash
		input       string
		expect      string
		shouldError bool
	}{
		{"MD5", MD5, "hello", "5d41402abc4b2a76b9719d911017c592", false},
		{"SHA1", SHA1, "hello", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", false},
		{"SHA256", SHA256, "hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", false},
		{"SHA512", SHA512, "hello", "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043", false},
		{"CRC32", CRC32, "hello", "3610a686", false},
		{"FNV32", FNV32, "hello", "b6fa7167", false},
		{"FNV32A", FNV32A, "hello", "4f9f2cab", false},
		{"FNV64", FNV64, "hello", "7b495389bdbdd4c7", false},
		{"FNV64A", FNV64A, "hello", "a430d84680aabd0b", false},
		{"FNV128", FNV128, "hello", "f14b58486483d94f708038798c29697f", false},
		{"FNV128A", FNV128A, "hello", "e3e1efd54283d94f7081314b599d31b3", false},
		// 添加自定义哈希算法的测试用例
		{"CUSTOM", 0, "hello", "mock_custom_hash", false},
		{
			name:        "CUSTOM",
			input:       "test123",
			expect:      "mock_custom_hash",
			shouldError: false,
		},
		{
			name:        "CUSTOM2",
			input:       "hello",
			expect:      "custom_hash_2",
			shouldError: false,
		},
		{
			name:        "CUSTOM3",
			input:       "data",
			expect:      "custom3-4",
			shouldError: true,
		},

		// 异常case
		{
			name:        "NOT_EXIST",
			input:       "test",
			shouldError: true,
		},
		{
			name:        "CONFLICT",
			input:       "test",
			expect:      "version2",
			shouldError: false,
		},
		{
			name:        "INVALID!NAME",
			input:       "test",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.hash == 0 {
				tt.hash, err = ParseHash(tt.name)
				if tt.shouldError {
					if err == nil {
						t.Fatalf("ParseHash(%s) error = %v, wantErr %v", tt.name, err, tt.shouldError)
					}
					return
				}
			} else {

			}
			h := tt.hash.New()
			h.Write([]byte(tt.input))
			got := h.Sum(nil)
			gotHex := string(got)
			if !IsCustomHash(tt.hash) {
				gotHex = fmt.Sprintf("%x", got)
			}
			if gotHex != tt.expect {
				t.Errorf("%s() = %v, want %v", tt.name, gotHex, tt.expect)
			}
		})
	}
}
