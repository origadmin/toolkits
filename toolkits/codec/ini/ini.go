/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package ini

import (
	"bytes"
	"io"

	"gopkg.in/ini.v1"
)

type Decoder struct {
	dec *ini.File
	err error
}

func (d Decoder) Decode(obj any) error {
	if d.err != nil {
		return d.err
	}
	return d.dec.MapTo(obj)
}

func NewDecoder(r io.Reader) *Decoder {
	load, err := ini.Load(r)
	if err != nil {
		return &Decoder{err: err}
	}
	return &Decoder{dec: load}
}

type Encoder struct {
	w io.Writer
}

func (e Encoder) Encode(obj any) error {
	cfg := ini.Empty()
	if err := ini.ReflectFrom(cfg, obj); err != nil {
		return err
	}
	if _, err := cfg.WriteTo(e.w); err != nil {
		return err
	}
	return nil
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func Unmarshal(data []byte, v any) error {
	cfg, err := ini.Load(data)
	if err != nil {
		return err
	}
	dec := &Decoder{dec: cfg}
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}

func Marshal(v any) ([]byte, error) {
	cfg := ini.Empty()
	if err := ini.ReflectFrom(cfg, v); err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	enc := Encoder{w: buf}
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

// MarshalToString returns json string, and ignores error
func MarshalToString(v any) string {
	b, err := Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

// MustToString returns json string, or panic
func MustToString(v any) string {
	data, err := Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Decode decodes the given JSON string into v
func Decode(data string, v any) error {
	return Unmarshal([]byte(data), v)
}
