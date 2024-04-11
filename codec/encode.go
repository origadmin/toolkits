// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package codec is the codec package for encoding and decoding
package codec

import (
	"os"
	"path/filepath"

	"github.com/origadmin/toolkits/errors"

	"github.com/origadmin/toolkits/codec/json"
	"github.com/origadmin/toolkits/codec/toml"
	"github.com/origadmin/toolkits/codec/yaml"
)

const (
	ErrUnsupportedEncodeType = errors.StdError("codec: unsupported encode type")
)

// EncodeJSONFile Encodes the given JSON file
func EncodeJSONFile(name string, obj any) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer errorPanic(f.Close())
	err = json.NewEncoder(f).Encode(obj)
	if err != nil {
		return err
	}
	return err
}

// EncodeYAMLFile Encodes the given YAML file
func EncodeYAMLFile(name string, obj any) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer errorPanic(f.Close())
	err = yaml.NewEncoder(f).Encode(obj)
	if err != nil {
		return err
	}
	return err
}

// EncodeTOMLFile Encodes the given TOML file
func EncodeTOMLFile(name string, obj any) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer errorPanic(f.Close())
	err = toml.NewEncoder(f).Encode(obj)
	if err != nil {
		return err
	}
	return err
}

// EncodeFile Encodes the given file
func EncodeFile(name string, obj any) error {
	ext := filepath.Ext(name)
	if ext == "" || !isSupportedExt(ext) {
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
