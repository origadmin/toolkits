/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package codec

import (
	"io"
	"path/filepath"

	"github.com/origadmin/toolkits/codec/ini"
	"github.com/origadmin/toolkits/codec/json"
	"github.com/origadmin/toolkits/codec/toml"
	"github.com/origadmin/toolkits/codec/xml"
	"github.com/origadmin/toolkits/codec/yaml"
)

//go:generate go run -mod=mod golang.org/x/tools/cmd/stringer@latest -type=Type
type Type int

const (
	JSON Type = iota // json
	YAML             // yaml
	TOML             // toml
	XML
	INI
	UNKNOWN // unknown
	TypeMax = UNKNOWN
)

var (
	codecs = [TypeMax]Codec{
		JSON: json.Codec,
		YAML: yaml.Codec,
		TOML: toml.Codec,
		XML:  xml.Codec,
		INI:  ini.Codec,
	}
)

func (s Type) Marshal(v interface{}) ([]byte, error) {
	if s >= UNKNOWN {
		return nil, ErrUnsupportedEncodeType
	}
	return codecs[s].Marshal(v)
}

func (s Type) Unmarshal(data []byte, v interface{}) error {
	if s >= UNKNOWN {
		return ErrUnsupportedDecodeType
	}
	return codecs[s].Unmarshal(data, v)
	//return s.NewDecoder(bytes.NewReader(data)).Decode(v)
}

func (s Type) NewDecoder(r io.Reader) Decoder {
	switch s {
	case JSON:
		return json.NewDecoder(r)
	case YAML:
		return yaml.NewDecoder(r)
	case TOML:
		return toml.NewDecoder(r)
	case XML:
		return xml.NewDecoder(r)
	case INI:
		return ini.NewDecoder(r)
	default:
		return nil
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
	case XML:
		return xml.NewEncoder(w)
	case INI:
		return ini.NewEncoder(w)
	default:
		return nil
	}
}

func (s Type) Name() string {
	if s == UNKNOWN {
		return "unknown"
	}
	return codecs[s].Name()
	//return s.String()
}

func (s Type) IsSupported() bool {
	return s != UNKNOWN
}

func (s Type) Exts() []string {
	switch s {
	case JSON:
		return []string{".json"}
	case YAML:
		return []string{".yaml", ".yml"}
	case TOML:
		return []string{".toml"}
	case XML:
		return []string{".xml"}
	case INI:
		return []string{".ini"}
	default:
		return []string{}
	}
}

// TypeFromString returns the codec type from the string.
func TypeFromString(s string) Type {
	switch s {
	case "json":
		return JSON
	case "yaml", "yml":
		return YAML
	case "toml":
		return TOML
	case "xml":
		return XML
	case "ini":
		return INI
	default:
		return UNKNOWN
	}
}

// TypeFromExt returns the codec type from the file extension.
func TypeFromExt(ext string) Type {
	switch ext {
	case ".json":
		return JSON
	case ".yaml", ".yml":
		return YAML
	case ".toml":
		return TOML
	case ".xml":
		return XML
	case ".ini":
		return INI
	default:
		return UNKNOWN
	}
}

// TypeFromPath returns the codec type from the file path.
func TypeFromPath(path string) Type {
	return TypeFromExt(filepath.Ext(path))
}

// IsSupported returns true if the codec type is supported.
func IsSupported(s string) bool {
	return TypeFromString(s).IsSupported()
}
