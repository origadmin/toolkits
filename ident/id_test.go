// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package ident provides the helpers functions.
package ident

import (
	"strings"
	"testing"
)

func TestNewXID(t *testing.T) {
	t.Logf("xid: %s", strings.ToUpper(NewXID()))
	t.Logf("xid string size: %d", len(NewXID()))
}

func TestMustNewUUID(t *testing.T) {
	t.Logf("uuid: %s", strings.ToUpper(MustNewUUID()))
	t.Logf("uuid string size: %d", len(MustNewUUID()))
}

func TestMustNewULID(t *testing.T) {
	t.Logf("ulid: %s", strings.ToUpper(MustNewULID()))
	t.Logf("ulid string size: %d", len(MustNewULID()))
	MustNewULID()
}
