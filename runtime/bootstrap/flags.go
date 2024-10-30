package bootstrap

import (
	"os"
	"time"
)

type Flags struct {
	ID          string
	Version     string
	ServiceName string
	StartTime   time.Time
	Metadata    map[string]string
}

func (f Flags) ServiceID() string {
	return f.ID + "." + f.ServiceName
}

func DefaultFlags() Flags {
	id, _ := os.Hostname()
	return Flags{
		ID:        id,
		StartTime: time.Now(),
	}
}

func NewFlags(name string, version string) Flags {
	f := DefaultFlags()
	f.Version = version
	f.ServiceName = name
	return f
}
