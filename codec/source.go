package codec

import (
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

type source struct {
	source config.Source
}

func (s *source) Load() ([]*config.KeyValue, error) {
	return s.source.Load()
}

func (s *source) Watch() (config.Watcher, error) {
	return s.source.Watch()
}

func NewSource(path string) config.Source {
	return &source{
		source: file.NewSource(path),
	}
}

func SourceDecoder(src *config.KeyValue, target map[string]interface{}) error {
	if src.Format == "" {
		// expand key "aaa.bbb" into map[aaa]map[bbb]interface{}
		keys := strings.Split(src.Key, ".")
		for i, k := range keys {
			if i == len(keys)-1 {
				target[k] = src.Value
			} else {
				sub := make(map[string]interface{})
				target[k] = sub
				target = sub
			}
		}
		return nil
	}
	if codec := SupportTypeFromString(src.Format); codec != UNKNOWN {
		return codec.Unmarshal(src.Value, &target)
	}
	return fmt.Errorf("unsupported key: %s format: %s", src.Key, src.Format)
}

var _ config.Source = (*source)(nil)
