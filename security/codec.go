package security

type Encoder interface {
	Encode(args ...any) (string, error)
}

type Decoder interface {
	Decode(string) (string, bool)
}
