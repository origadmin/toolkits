package codec

import (
	"bytes"

	"github.com/origadmin/toolkits/codec/json"
	"github.com/origadmin/toolkits/codec/toml"
	"github.com/origadmin/toolkits/codec/yaml"
	"github.com/origadmin/toolkits/io"
)

//go:generate go run -mod=mod golang.org/x/tools/cmd/stringer@latest -type=Type
type Type int

const (
	JSON    Type = iota // json
	YAML                // yaml
	TOML                // toml
	UNKNOWN             // unknown
)

func (s Type) Marshal(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := s.NewEncoder(buf).Encode(v)
	return buf.Bytes(), err
}

func (s Type) Unmarshal(data []byte, v interface{}) error {
	return s.NewDecoder(bytes.NewReader(data)).Decode(v)
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

func (s Type) Name() string {
	return s.String()
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
	default:
		return UNKNOWN
	}
}

// TypeFromExt returns the codec type from the file extension.
func TypeFromExt(ext string) Type {
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
