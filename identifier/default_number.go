/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package identifier

import (
	"fmt"
	"sync"
	"time"
)

// --- Built-in Default Provider for int64 (Fallback) ---

const (
	snowflakeEpoch     = 1704067200000 // 2024-01-01 00:00:00 UTC in milliseconds
	machineIDBits    = 10
	sequenceBits     = 12
	maxMachineID     = -1 ^ (-1 << machineIDBits)
	maxSequence      = -1 ^ (-1 << sequenceBits)
	timestampShift   = machineIDBits + sequenceBits
	machineIDShift   = sequenceBits
)

// defaultSnowflakeProvider is a minimal, dependency-free Snowflake generator.
// It is used as a fallback when no other provider for "snowflake" is registered.
// Note: The machine ID is hardcoded to 1. For production use in a distributed
// environment, a proper snowflake implementation that assigns unique machine IDs
// should be registered.
type defaultSnowflakeProvider struct {
	mu        sync.Mutex
	lastStamp int64
	sequence  int64
	machineID int64
}

func (p *defaultSnowflakeProvider) Name() string { return "snowflake" }
func (p *defaultSnowflakeProvider) Size() int    { return 64 }
func (p *defaultSnowflakeProvider) AsString() Generator[string] { return nil }
func (p *defaultSnowflakeProvider) AsNumber() Generator[int64] { return p }

func (p *defaultSnowflakeProvider) Generate() int64 {
	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now().UnixMilli()

	if now < p.lastStamp {
		// Clock moved backwards, refuse to generate ID.
		panic(fmt.Sprintf("identifier: clock moved backwards. Refusing to generate id for %d milliseconds", p.lastStamp-now))
	}

	if p.lastStamp == now {
		p.sequence = (p.sequence + 1) & maxSequence
		if p.sequence == 0 {
			// Sequence overflow, wait for next millisecond.
			for now <= p.lastStamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		p.sequence = 0
	}

	p.lastStamp = now

	return ((now - snowflakeEpoch) << timestampShift) |
		(p.machineID << machineIDShift) |
		p.sequence
}

func (p *defaultSnowflakeProvider) Validate(id int64) bool {
	// Basic validation: check if the timestamp part is not in the future.
	// A full validation is more complex and provider-specific.
	timestamp := (id >> timestampShift) + snowflakeEpoch
	return timestamp <= time.Now().UnixMilli()
}

// builtinNumber is the singleton instance of our built-in fallback for int64.
var builtinNumber Provider = &defaultSnowflakeProvider{machineID: 1}
