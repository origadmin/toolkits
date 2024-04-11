package errors_test

import (
	"testing"

	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/errors/http"
)

func TestFind(t *testing.T) {
	type T = http.Error
	// Test case when err is of type T
	errOfTypeT := http.New("error id", 0, "error message")
	if result := errors.Find[*T](errOfTypeT); !errors.Is(result, errOfTypeT) {
		t.Errorf("Expected %v, but got %v", errOfTypeT, result)
	}

	// Test case when err is not of type T
	errNotOfTypeT := errors.New("error not of type T")
	if result := errors.Find[*T](errNotOfTypeT); result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}

	// Test case when err is nil
	if result := errors.Find[*T](nil); result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}

	errOfTypeT = http.New("error id", 0, "error message")
	errOfTypeT = errors.Wrapf(errOfTypeT, "wrap error")
	if result := errors.Find[*T](errOfTypeT); !errors.Is(result, errOfTypeT) {
		t.Errorf("Expected %v, but got %v", errOfTypeT, result)
	}

	errOfTypeT = errors.New("error message")
	errOfTypeT = errors.Wrapf(errOfTypeT, "wrap error")
	if result := errors.Find[errors.StdError](errOfTypeT); !errors.Is(result, errOfTypeT) {
		t.Errorf("Expected %v, but got %v", errOfTypeT, result)
	}

	errOfTypeT = errors.NewErr("error message")
	errOfTypeT = errors.Wrapf(errOfTypeT, "wrap error")
	if result := errors.Find[errors.StdError](errOfTypeT); result != nil {
		t.Errorf("Expected %v, but got %v", errOfTypeT, result)
	}

	errOfTypeT = errors.NewErr("error message")
	if result := errors.New("error message"); errors.Is(errOfTypeT, result) {
		t.Errorf("Expected %v, but got %v", errOfTypeT, result)
	}

	if result := errors.Find[error](errOfTypeT); result != nil {
		t.Errorf("Expected %v, but got %v", errOfTypeT, result)
	}

}
