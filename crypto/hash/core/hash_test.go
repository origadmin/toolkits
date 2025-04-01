package core

import (
	"fmt"
	"hash"
	"testing"
)

// 定义一个自定义哈希算法的类型
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

// 实现自定义哈希算法的 Sum 方法
func (c *CustomHash) Sum(b []byte) []byte {
	return []byte("mock_custom_hash")
}

// 注册自定义哈希算法
func init() {
	RegisterHashFunc("CUSTOM", func() hash.Hash {
		return &CustomHash{}
	})
}

func TestHashAlgorithms(t *testing.T) {
	tests := []struct {
		name   string
		hash   Hash
		input  string
		expect string
	}{
		{"MD5", MD5, "hello", "5d41402abc4b2a76b9719d911017c592"},
		{"SHA1", SHA1, "hello", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"},
		{"SHA256", SHA256, "hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"SHA512", SHA512, "hello", "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"},
		{"CRC32", CRC32, "hello", "3610a686"},
		{"FNV32", FNV32, "hello", "b6fa7167"},
		{"FNV32A", FNV32A, "hello", "4f9f2cab"},
		{"FNV64", FNV64, "hello", "7b495389bdbdd4c7"},
		{"FNV64A", FNV64A, "hello", "a430d84680aabd0b"},
		{"FNV128", FNV128, "hello", "f14b58486483d94f708038798c29697f"},
		{"FNV128A", FNV128A, "hello", "e3e1efd54283d94f7081314b599d31b3"},
		// 添加自定义哈希算法的测试用例
		{"CUSTOM", 0, "hello", fmt.Sprintf("%x", []byte("mock_custom_hash"))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.hash == 0 {
				tt.hash, err = ParseHash("CUSTOM")
				if err != nil {
					t.Fatalf("ParseHash(%s) error = %v", tt.name, err)
				}
			} else {

			}
			h := tt.hash.New()
			h.Write([]byte(tt.input))
			got := h.Sum(nil)
			if gotHex := fmt.Sprintf("%x", got); gotHex != tt.expect {
				t.Errorf("%s() = %v, want %v", tt.name, gotHex, tt.expect)
			}
		})
	}
}
