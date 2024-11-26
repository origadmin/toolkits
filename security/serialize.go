/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package security

import (
	"time"
)

type TokenSerializer interface {
	Generate(subject string, expires ...time.Duration) Token
	Parse(token string) (Token, error)
}
