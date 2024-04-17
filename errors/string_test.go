package errors

import (
	er "errors"
	"testing"
)

func TestAs(t *testing.T) {
	var err error = String("errors.rpc.test")
	var target String
	match := er.As(err, &target)
	if !match {
		t.Fatalf("%v should convert to *Error", err)
	}

	err = er.New(err.Error())
	target = ""
	match = er.As(err, &target)
	if match || target != "" {
		t.Fatalf("%v should not convert to *Error", err)
	}
}

func TestIs(t *testing.T) {
	var err error = String("errors.rpc.test")
	var target = String("errors.rpc.test")
	match := er.Is(err, &target)
	if !match {
		t.Fatalf("%v should convert to String", err)
	}

	err = er.New(err.Error())
	target = "errors.rpc.test"
	match = er.Is(err, &target)
	if match {
		t.Fatalf("%v should not equal to String", err)
	}
}
