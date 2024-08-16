// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package codec is the codec package for encoding and decoding
package codec

import (
	"io"
	"os"
	"path/filepath"

	"github.com/origadmin/toolkits/errors"

	"github.com/origadmin/toolkits/codec/json"
	"github.com/origadmin/toolkits/codec/toml"
	"github.com/origadmin/toolkits/codec/yaml"
)

const (
	ErrUnsupportedEncodeType = errors.String("codec: unsupported encode type")
)

// Encoder interface
type Encoder interface {
	Encode(obj any) error
}

// EncodeJSONFile Encodes the given JSON file
func EncodeJSONFile(name string, obj any) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer func() {
		errorPanic(f.Close())
	}()
	return json.NewEncoder(f).Encode(obj)
}

// EncodeYAMLFile Encodes the given YAML file
func EncodeYAMLFile(name string, obj any) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer func() {
		errorPanic(f.Close())
	}()
	return yaml.NewEncoder(f).Encode(obj)
}

// EncodeTOMLFile Encodes the given TOML file
func EncodeTOMLFile(name string, obj any) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer func() {
		errorPanic(f.Close())
	}()
	return toml.NewEncoder(f).Encode(obj)
}

// EncodeFile Encodes the given file
func EncodeFile(name string, obj any) error {
	ext := filepath.Ext(name)
	if ext == "" || !isSupported(ext) {
		return ErrUnsupportedEncodeType
	}
	switch filepath.Ext(name) {
	case ".json":
		return EncodeJSONFile(name, obj)
	case ".yaml", ".yml":
		return EncodeYAMLFile(name, obj)
	case ".toml":
		return EncodeTOMLFile(name, obj)
	default:
		return ErrUnsupportedEncodeType
	}
}

func Encode(w io.Writer, obj any, st Type) error {
	switch st {
	case JSON:
		return json.NewEncoder(w).Encode(obj)
	case YAML:
		return yaml.NewEncoder(w).Encode(obj)
	case TOML:
		return toml.NewEncoder(w).Encode(obj)
	default:
		return ErrUnsupportedEncodeType
	}
}
