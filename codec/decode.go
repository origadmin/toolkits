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
	ErrUnsupportedDecodeType = errors.StdError("codec: unsupported decode type")
)

// DecodeJSONFile Decodes the given JSON file
func DecodeJSONFile(name string, obj any) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(obj)
}

// DecodeTOMLFile Decodes the given TOML file
func DecodeTOMLFile(name string, obj any) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = toml.NewDecoder(f).Decode(obj)
	if err != nil {
		return err
	}
	return err
}

// DecodeYAMLFile Decodes the given YAML file
func DecodeYAMLFile(name string, obj any) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	return yaml.NewDecoder(f).Decode(obj)
}

// DecodeFile Decodes the given file
func DecodeFile(name string, obj any) error {
	ext := filepath.Ext(name)
	if ext == "" || !isSupportedExt(ext) {
		return ErrUnsupportedDecodeType
	}

	switch filepath.Ext(name) {
	case ".json":
		return DecodeJSONFile(name, obj)
	case ".yaml", ".yml":
		return DecodeYAMLFile(name, obj)
	case ".toml":
		return DecodeTOMLFile(name, obj)
	default:
		return ErrUnsupportedDecodeType
	}
}
