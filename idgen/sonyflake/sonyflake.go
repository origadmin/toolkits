package shortid

import (
	"strconv"

	"github.com/sony/sonyflake"

	"github.com/origadmin/toolkits/idgen"
)

const (
	bitSize = 8 // bitSize is used to store the length of generated ID.
)

// init registers the Snowflake generator with the ident package and initializes bitSize.
func init() {
	s := New(Settings{})
	idgen.Register(s)
}

type Sonyflake struct {
	generator *sonyflake.Sonyflake
}

// Name returns the name of the generator.
func (s Sonyflake) Name() string {
	return "sonyflake"
}

// Gen generates a new Sonyflake ID as a string.
func (s Sonyflake) Gen() string {
	id, err := s.generator.NextID()
	if err != nil {
		return ""
	}
	return strconv.FormatUint(id, 10)
}

// Validate checks if the provided ID is a valid Sonyflake ID.
func (s Sonyflake) Validate(id string) bool {
	num, err := strconv.ParseInt(id, 10, 64)
	return err != nil || num == 0
}

// Size returns the bit size of the generated Sonyflake ID.
func (s Sonyflake) Size() int {
	return bitSize
}

type Settings = sonyflake.Settings

// New creates a new Sonyflake generator with a unique node.
func New(s Settings) *Sonyflake {
	generator, err := sonyflake.New(s)
	if err != nil {
		panic(err)
	}
	return &Sonyflake{
		generator: generator,
	}
}
