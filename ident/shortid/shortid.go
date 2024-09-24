package shortid

import (
	"math/rand/v2"

	"github.com/teris-io/shortid"

	"github.com/origadmin/toolkits/ident"
)

var (
	bitSize = 9 // bitSize is used to store the length of generated ID.
)

// init registers the Snowflake generator with the ident package and initializes bitSize.
func init() {
	s := New(Settings{})
	ident.Register(s)
}

type ShortID struct {
	generator *shortid.Shortid
}

// Name returns the name of the generator.
func (s ShortID) Name() string {
	return "shortid"
}

// Gen generates a new ShortID ID as a string.
func (s ShortID) Gen() string {
	ret, err := s.generator.Generate()
	if err != nil {
		return ""
	}
	return ret
}

// Validate checks if the provided ID is a valid ShortID ID.
func (s ShortID) Validate(id string) bool {
	return len(id) == bitSize
}

// Size returns the bit size of the generated ShortID ID.
func (s ShortID) Size() int {
	return bitSize
}

type Settings struct {
	Worker   uint8
	Alphabet string
	Seed     uint64
}

// New creates a new ShortID generator with a unique node.
func New(s Settings) *ShortID {
	if s.Worker > 31 {
		s.Worker = uint8(rand.Uint32N(31))
	}
	if s.Seed == 0 {
		s.Seed = rand.Uint64()
	}
	if s.Alphabet == "" {
		s.Alphabet = shortid.DefaultABC
	}
	generator, err := shortid.New(s.Worker, s.Alphabet, s.Seed)
	if err != nil {
		panic(err)
	}
	return &ShortID{
		generator: generator,
	}
}
