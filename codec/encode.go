// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package codec is the codec package for encoding and decoding
package codec

import (
	"encoding/xml"
	"io"
	"os"
	"path/filepath"

	"github.com/origadmin/toolkits/codec/ini"
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
	defer f.Close()
	return json.NewEncoder(f).Encode(obj)
}

// EncodeYAMLFile Encodes the given YAML file
func EncodeYAMLFile(name string, obj any) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewEncoder(f).Encode(obj)
}

// EncodeTOMLFile Encodes the given TOML file
func EncodeTOMLFile(name string, obj any) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(obj)
}

// EncodeXMLFile Encodes the given XML file
func EncodeXMLFile(name string, obj any) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	return xml.NewEncoder(f).Encode(obj)
}

// EncodeINIFile Encodes the given INI file
func EncodeINIFile(name string, obj any) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	return ini.NewEncoder(f).Encode(obj)
}

// EncodeToFile Encodes the given file
func EncodeToFile(name string, obj any) error {
	typo := TypeFromExt(filepath.Ext(name))
	if !typo.IsSupported() {
		return ErrUnsupportedEncodeType
	}
	switch filepath.Ext(name) {
	case ".json":
		return EncodeJSONFile(name, obj)
	case ".yaml", ".yml":
		return EncodeYAMLFile(name, obj)
	case ".toml":
		return EncodeTOMLFile(name, obj)
	case ".xml":
		return EncodeXMLFile(name, obj)
	case ".ini":
		return EncodeINIFile(name, obj)
	default:
		return ErrUnsupportedEncodeType
	}
}

// Encode Encodes the given object with ext name
func Encode(w io.Writer, obj any, st Type) error {
	switch st {
	case JSON:
		return json.NewEncoder(w).Encode(obj)
	case YAML:
		return yaml.NewEncoder(w).Encode(obj)
	case TOML:
		return toml.NewEncoder(w).Encode(obj)
	case XML:
		return xml.NewEncoder(w).Encode(obj)
	case INI:
		return ini.NewEncoder(w).Encode(obj)
	default:
		return ErrUnsupportedEncodeType
	}
}
