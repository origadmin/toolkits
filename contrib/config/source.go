package config

import (
	"strings"

	"github.com/go-kratos/kratos/v2/config"

	"github.com/origadmin/toolkits/codec"
	"github.com/origadmin/toolkits/errors"
)

type (
	// Source is config source.
	Source = config.Source
	// Config is config.
	Config = config.Config
	// Watcher is config watcher.
	Watcher = config.Watcher
	// Option is config option.
	Option = config.Option
	// Decoder is config decoder.
	Decoder = config.Decoder
	// Resolver resolve placeholder in config.
	Resolver = config.Resolver
	// Merge is config merge func.
	Merge = config.Merge
	// KeyValue is config key value.
	KeyValue = config.KeyValue
	// Value is config value.
	Value = config.Value
	// Reader is config reader.
	Reader = config.Reader
)

var (
	// WithSource with config source.
	WithSource = config.WithSource

	// WithDecoder with config decoder.
	// DefaultDecoder behavior:
	// If KeyValue.Format is non-empty, then KeyValue.Value will be deserialized into map[string]interface{}
	// and stored in the config cache(map[string]interface{})
	// if KeyValue.Format is empty,{KeyValue.Key : KeyValue.Value} will be stored in config cache(map[string]interface{})
	WithDecoder = config.WithDecoder

	// WithResolveActualTypes with config resolver.
	// bool input will enable conversion of config to data types
	WithResolveActualTypes = config.WithResolveActualTypes

	// WithResolver with config resolver.
	WithResolver = config.WithResolver

	// WithMergeFunc with config merge func.
	WithMergeFunc = config.WithMergeFunc
)

func New(opts ...Option) Config {
	opt := config.WithDecoder(SourceDecoder)
	opts = append(opts, opt)
	return config.New(opts...)
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
	if codec := codec.TypeFromString(src.Format); codec.IsSupported() {
		return codec.Unmarshal(src.Value, &target)
	}
	return errors.Errorf("unsupported key: %s format: %s", src.Key, src.Format)
}
