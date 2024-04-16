package rpc

import (
	er "errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadRequest(t *testing.T) {
	// Testing with a non-empty ID
	err1 := BadRequest("123", "Error message %s", "object1")
	if err1 == nil {
		t.Errorf("Expected an error to be returned with non-empty ID, but got nil")
	}
	assert.EqualError(t, err1, `{"id":"123","code":400,"detail":"Error message object1"}`)

	// Testing with an empty ID
	err2 := BadRequest("", "Error message %s", "object2")
	if err2 == nil {
		t.Errorf("Expected an error to be returned with empty ID, but got nil")
	}
	assert.EqualError(t, err2, `{"id":"http.response.status.bad_request","code":400,"detail":"Error message object2"}`)

	// Testing with different formats and objects provided
	err3 := BadRequest("456", "Error message without objects")
	if err3 == nil {
		t.Errorf("Expected an error to be returned with format but no objects, but got nil")
	}
	assert.EqualError(t, err3, `{"id":"456","code":400,"detail":"Error message without objects"}`)
	err4 := BadRequest("789", "Error message with multiple objects %s %s", "object3", "object4")
	if err4 == nil {
		t.Errorf("Expected an error to be returned with multiple objects, but got nil")
	}
	assert.EqualError(t, err4, `{"id":"789","code":400,"detail":"Error message with multiple objects object3 object4"}`)
}

func TestFromError(t *testing.T) {
	err := NotFound("errors.rpc.test", "%s", "example")
	merr := FromError(err)
	if merr.Id != "errors.rpc.test" || merr.Code != 404 {
		t.Fatalf("invalid conversation %v != %v", err, merr)
	}
	err = er.New(err.Error())
	merr = FromError(err)
	if merr.Id != "errors.rpc.test" || merr.Code != 404 {
		t.Fatalf("invalid conversation %v != %v", err, merr)
	}
	merr = FromError(nil)
	if merr != nil {
		t.Fatalf("%v should be nil", merr)
	}
}

func TestEqual(t *testing.T) {
	err1 := NotFound("id1", "msg1")
	err2 := NotFound("id2", "msg2")

	if !Equal(err1, err2) {
		t.Fatal("errors must be equal")
	}

	err1 = NotFound("id1", "msg1")
	err2 = InternalServerError("id2", "msg2")

	if Equal(err1, err2) {
		t.Fatal("errors must be not equal")
	}

	err3 := er.New("my test err")
	if Equal(err1, err3) {
		t.Fatal("errors must be not equal")
	}

}

func TestErrors(t *testing.T) {
	testData := []*Error{
		{
			Id:     "test",
			Code:   500,
			Detail: "Internal server error",
		},
	}

	for _, e := range testData {
		ne := New(e.Id, e.Code, e.Detail)

		if e.Error() != ne.Error() {
			t.Fatalf("Expected %s got %s", e.Error(), ne.Error())
		}

		pe := Parse(ne.Error())

		if pe == nil {
			t.Fatalf("Expected error got nil %v", pe)
		}

		if pe.Id != e.Id {
			t.Fatalf("Expected %s got %s", e.Id, pe.Id)
		}

		if pe.Detail != e.Detail {
			t.Fatalf("Expected %s got %s", e.Detail, pe.Detail)
		}

		if pe.Code != e.Code {
			t.Fatalf("Expected %d got %d", e.Code, pe.Code)
		}
	}
}

func TestAs(t *testing.T) {
	err := NotFound("errors.rpc.test", "%s", "example")
	var target *Error
	match := er.As(err, &target)
	if !match {
		t.Fatalf("%v should convert to *Error", err)
	}
	if target.Id != "errors.rpc.test" || target.Code != 404 || target.Detail != "example" {
		t.Fatalf("invalid conversation %v != %v", err, target)
	}
	err = er.New(err.Error())
	target = nil
	match = er.As(err, &target)
	if match || target != nil {
		t.Fatalf("%v should not convert to *Error", err)
	}
}

func TestAppend(t *testing.T) {
	mError := NewMultiError()
	testData := []*Error{
		{
			Id:     "test1",
			Code:   500,
			Detail: "Internal server error",
		},
		{
			Id:     "test2",
			Code:   400,
			Detail: "Bad Request",
		},
		{
			Id:     "test3",
			Code:   404,
			Detail: "Not Found",
		},
	}

	for _, e := range testData {
		mError.Append(&Error{
			Id:     e.Id,
			Code:   e.Code,
			Detail: e.Detail,
		})
	}

	if len(mError.Errors) != 3 {
		t.Fatalf("Expected 3 got %v", len(mError.Errors))
	}
}

func TestHasErrors(t *testing.T) {
	mError := NewMultiError()
	testData := []*Error{
		{
			Id:     "test1",
			Code:   500,
			Detail: "Internal server error",
		},
		{
			Id:     "test2",
			Code:   400,
			Detail: "Bad Request",
		},
		{
			Id:     "test3",
			Code:   404,
			Detail: "Not Found",
		},
	}

	if mError.HasErrors() {
		t.Fatal("Expected no error")
	}

	for _, e := range testData {
		mError.Errors = append(mError.Errors, &Error{
			Id:     e.Id,
			Code:   e.Code,
			Detail: e.Detail,
		})
	}

	if !mError.HasErrors() {
		t.Fatal("Expected errors")
	}
}
