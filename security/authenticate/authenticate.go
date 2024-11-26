/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package authenticate

type Authenticator interface {
	Scheme() AuthScheme
	Credentials() string
	Extra() []string
	Encode(args ...any) (string, error)
}
