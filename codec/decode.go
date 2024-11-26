/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package codec is the codec package for encoding and decoding
package codec

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/origadmin/toolkits/codec/ini"
	"github.com/origadmin/toolkits/codec/json"
	"github.com/origadmin/toolkits/codec/toml"
	"github.com/origadmin/toolkits/codec/xml"
	"github.com/origadmin/toolkits/codec/yaml"
)

var (
	ErrUnsupportedDecodeType = errors.New("codec: unsupported decode type")
)

// Decoder interface
type Decoder interface {
	Decode(any) error
}

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

	err = toml.NewDecoder(f).Decode(obj)
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

// DecodeXMLFile Decodes the given XML file
func DecodeXMLFile(name string, obj any) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return xml.NewDecoder(f).Decode(obj)
}

// DecodeINIFile Decodes the given INI file
func DecodeINIFile(name string, obj any) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return ini.NewDecoder(f).Decode(obj)
}

// DecodeFromFile Decodes the given file
func DecodeFromFile(name string, obj any) error {
	dec := TypeFromExt(filepath.Ext(name))
	if !dec.IsSupported() {
		return ErrUnsupportedDecodeType
	}

	switch filepath.Ext(name) {
	case ".json":
		return DecodeJSONFile(name, obj)
	case ".yaml", ".yml":
		return DecodeYAMLFile(name, obj)
	case ".toml":
		return DecodeTOMLFile(name, obj)
	case ".xml":
		return DecodeXMLFile(name, obj)
	case ".ini":
		return DecodeINIFile(name, obj)
	default:
		return ErrUnsupportedDecodeType
	}
}

// Decode Decodes the given reader with ext name into obj
func Decode(rd io.Reader, obj any, ext string) error {
	switch ext {
	case ".json":
		return json.NewDecoder(rd).Decode(obj)
	case ".yaml", ".yml":
		return yaml.NewDecoder(rd).Decode(obj)
	case ".toml":
		return toml.NewDecoder(rd).Decode(obj)
	case ".ini":
		return ini.NewDecoder(rd).Decode(obj)
	default:
		return ErrUnsupportedDecodeType
	}
}
