/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package errors

import (
	stderr "errors"
	"testing"
)

func TestStringAs(t *testing.T) {
	var err error = String("errors.rpc.test")
	var target String
	match := stderr.As(err, &target)
	if !match {
		t.Fatalf("%v should convert to *Error", err)
	}

	err = stderr.New(err.Error())
	target = ""
	match = stderr.As(err, &target)
	if match || target != "" {
		t.Fatalf("%v should not convert to *Error", err)
	}
}

func TestStringIs(t *testing.T) {
	var err error = String("errors.rpc.test")
	var target = String("errors.rpc.test")
	match := stderr.Is(err, target)
	if !match {
		t.Fatalf("%v should convert to String", err)
	}

	err = stderr.New(err.Error())
	target = "errors.rpc.test"
	match = stderr.Is(err, &target)
	if match {
		t.Fatalf("%v should not equal to String", err)
	}
}
