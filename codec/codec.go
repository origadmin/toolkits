// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package codec is the codec package for encoding and decoding
package codec

// supportedExts support extensions of json, yaml, yml, toml
var supportedExts = []string{".json", ".yaml", ".yml", ".toml"}

// isSupportedExt 检查文件扩展名是否支持。
func isSupportedExt(ext string) bool {
	for _, supportedExt := range supportedExts {
		if ext == supportedExt {
			return true
		}
	}
	return false
}

// MustBytes returns bytes
func MustBytes(data []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return data
}

// MustToString returns string
func MustToString(data []byte, err error) string {
	if err != nil {
		panic(err)
	}
	return string(data)
}

// MustString returns string
func MustString(data string, err error) string {
	if err != nil {
		panic(err)
	}
	return data
}
