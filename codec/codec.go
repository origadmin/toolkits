// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package codec is the codec package for encoding and decoding
package codec

import (
	"bytes"
	"io"

	"github.com/BurntSushi/toml"
	ggb "github.com/goexts/ggb"

	"github.com/origadmin/toolkits/codec/json"
	"github.com/origadmin/toolkits/codec/yaml"
)

type Type int

func (s Type) Marshal(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := s.NewEncoder(buf).Encode(v)
	return buf.Bytes(), err
}

func (s Type) Unmarshal(data []byte, v interface{}) error {
	return s.NewDecoder(bytes.NewReader(data)).Decode(v)
}

func (s Type) Name() string {
	return s.String()
}

const (
	JSON Type = iota
	YAML
	TOML // toml
	UNKNOWN
)

func (s Type) String() string {
	switch s {
	case JSON:
		return "json"
	case YAML:
		return "yaml"
	case TOML:
		return "toml"
	default:
		return "unknown"
	}
}

func (s Type) Exts() []string {
	switch s {
	case JSON:
		return []string{".json"}
	case YAML:
		return []string{".yaml", ".yml"}
	case TOML:
		return []string{".toml"}
	default:
		return []string{}
	}
}

func (s Type) NewEncoder(w io.Writer) Encoder {
	switch s {
	case JSON:
		return json.NewEncoder(w)
	case YAML:
		return yaml.NewEncoder(w)
	case TOML:
		return toml.NewEncoder(w)
	default:
		return nil
	}
}

type tomlDecoder struct {
	dec *toml.Decoder
}

func (t tomlDecoder) Decode(obj any) error {
	return ggb.OrNil(t.dec.Decode(obj))
}

func (s Type) NewDecoder(r io.Reader) Decoder {
	switch s {
	case JSON:
		return json.NewDecoder(r)
	case YAML:
		return yaml.NewDecoder(r)
	case TOML:
		return &tomlDecoder{
			dec: toml.NewDecoder(r),
		}
	default:
		return nil
	}
}

// SupportTypeFromString returns the codec type from the string.
func SupportTypeFromString(s string) Type {
	switch s {
	case "json":
		return JSON
	case "yaml", "yml":
		return YAML
	case "toml":
		return TOML
	default:
		return UNKNOWN
	}
}

// SupportTypeFromExt returns the codec type from the file extension.
func SupportTypeFromExt(ext string) Type {
	if ext == "" {
		return UNKNOWN
	}
	switch ext {
	case ".json":
		return JSON
	case ".yaml", ".yml":
		return YAML
	case ".toml":
		return TOML
	default:
		return UNKNOWN
	}
}

// isSupported checks whether the file extension is supported。
func isSupported(ext string) bool {
	return SupportTypeFromExt(ext) != UNKNOWN
}

func errorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// MustBytes returns bytes
func MustBytes(data []byte, err error) []byte {
	errorPanic(err)
	return data
}

// MustToString returns string
func MustToString(data []byte, err error) string {
	errorPanic(err)
	return string(data)
}

// MustString returns string
func MustString(data string, err error) string {
	errorPanic(err)
	return data
}
