/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mail provides a mail client
package mail

import (
	"github.com/origadmin/toolkits/context"
)

// Sender is an interface for sending email
type Sender interface {
	// Send sends email to the given recipients.
	Send(ctx context.Context, to []string, cc []string, bcc []string, subject string, body string, file ...string) error
	// SendTo sends email to the given recipients.
	SendTo(ctx context.Context, to []string, subject string, body string, file ...string) error
}
