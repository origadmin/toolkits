package http_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/errors/http"
)

func TestBadRequest(t *testing.T) {
	// Testing with a non-empty ID
	err1 := http.BadRequest("123", "Error message %s", "object1")
	if err1 == nil {
		t.Errorf("Expected an error to be returned with non-empty ID, but got nil")
	}
	assert.EqualError(t, err1, `{"id":"123","code":400,"detail":"Error message object1"}`)

	// Testing with an empty ID
	err2 := http.BadRequest("", "Error message %s", "object2")
	if err2 == nil {
		t.Errorf("Expected an error to be returned with empty ID, but got nil")
	}
	assert.EqualError(t, err2, `{"id":"http.response.status.bad_request","code":400,"detail":"Error message object2"}`)

	// Testing with different formats and objects provided
	err3 := http.BadRequest("456", "Error message without objects")
	if err3 == nil {
		t.Errorf("Expected an error to be returned with format but no objects, but got nil")
	}
	assert.EqualError(t, err3, `{"id":"456","code":400,"detail":"Error message without objects"}`)
	err4 := http.BadRequest("789", "Error message with multiple objects %s %s", "object3", "object4")
	if err4 == nil {
		t.Errorf("Expected an error to be returned with multiple objects, but got nil")
	}
	assert.EqualError(t, err4, `{"id":"789","code":400,"detail":"Error message with multiple objects object3 object4"}`)
}
