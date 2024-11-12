package shortid

import (
	"cmp"
	"math/rand/v2"

	"github.com/teris-io/shortid"

	"github.com/origadmin/toolkits/idgen"
)

var (
	bitSize = 9 // bitSize is used to store the length of generated ID.
)

// init registers the Snowflake generator with the ident package and initializes bitSize.
func init() {
	s := New()
	idgen.Register(s)
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

type Setting struct {
	Worker   uint8
	Alphabet string
	Seed     uint64
}

// New creates a new ShortID generator with a unique node.
func New(ss ...*Setting) *ShortID {
	ss = append(ss, &Setting{})
	o := cmp.Or(ss...)
	if o.Worker > 31 {
		o.Worker = uint8(rand.Uint32N(31))
	}
	if o.Seed == 0 {
		o.Seed = rand.Uint64()
	}
	if o.Alphabet == "" {
		o.Alphabet = shortid.DefaultABC
	}
	generator, err := shortid.New(o.Worker, o.Alphabet, o.Seed)
	if err != nil {
		panic(err)
	}
	return &ShortID{
		generator: generator,
	}
}
