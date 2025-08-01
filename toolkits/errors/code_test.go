/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package errors

import (
	stderr "errors"
	"testing"
)

func TestCodeAs(t *testing.T) {
	var err error = Code(0)
	var target Code
	match := stderr.As(err, &target)
	if !match {
		t.Fatalf("%v should convert to *Error", err)
	}

	err = stderr.New(err.Error())
	target = 0
	match = stderr.As(err, &target)
	if match || target != 0 {
		t.Fatalf("%v should not convert to *Error", err)
	}
}

func TestCodeIs(t *testing.T) {
	var err error = Code(0)
	var target = Code(0)
	match := stderr.Is(err, target)
	if !match {
		t.Fatalf("%v should convert to Code", err)
	}

	err = stderr.New(err.Error())
	target = 0
	match = stderr.Is(err, &target)
	if match {
		t.Fatalf("%v should not equal to Code", err)
	}
}
