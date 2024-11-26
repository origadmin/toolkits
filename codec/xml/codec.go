/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package xml

var (
	Codec = codec{}
)

type codec struct{}

func (c codec) Marshal(v interface{}) ([]byte, error) {
	return Marshal(v)
}

func (c codec) Unmarshal(data []byte, v interface{}) error {
	return Unmarshal(data, v)
}

func (c codec) Name() string {
	return "xml"
}
