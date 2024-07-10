package http

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/origadmin/toolkits/errors"
)

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

	err3 := errors.New("my test err")
	if Equal(err1, err3) {
		t.Fatal("errors must be not equal")
	}

}

func TestErrors(t *testing.T) {
	testData := []*Error{
		{
			ID:     "test",
			Code:   500,
			Detail: "Internal server error",
		},
	}

	for _, e := range testData {
		ne := New(e.ID, e.Code, e.Detail)

		if e.Error() != ne.Error() {
			t.Fatalf("Expected %s got %s", e.Error(), ne.Error())
		}

		pe := Parse(ne.Error())

		if pe == nil {
			t.Fatalf("Expected error got nil %v", pe)
		}

		if pe.ID != e.ID {
			t.Fatalf("Expected %s got %s", e.ID, pe.ID)
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
	match := errors.As(err, &target)
	if !match {
		t.Fatalf("%v should convert to *Error", err)
	}
	if target.ID != "errors.rpc.test" || target.Code != 404 || target.Detail != "example" {
		t.Fatalf("invalid conversation %v != %v", err, target)
	}
	err = errors.New(err.Error())
	target = nil
	match = errors.As(err, &target)
	if match || target != nil {
		t.Fatalf("%v should not convert to *Error", err)
	}
}

func TestIs(t *testing.T) {
	var err = NotFound("errors.rpc.test", "%s", "example")
	var target = NotFound("errors.rpc.test", "%s", "example")
	match := errors.Is(err, target)
	if !match {
		t.Fatalf("%v should convert to String", err)
	}

	err = errors.New(err.Error())
	target = NotFound("errors.rpc.test", "%s", "example")
	match = errors.Is(err, target)
	if match {
		t.Fatalf("%v should not equal to String", err)
	}
}

func TestRegisterCode(t *testing.T) {
	// Test case 1: Register a new error code with a given ID
	code := int32(424)                   // Define a new error code
	id := "response.status.custom_error" // Define a new error ID
	RegisterCode(code, id)
	if got, want := ids[code], id; got != want {
		t.Errorf("RegisterCode(%d, %s) = %s; want %s", code, id, got, want)
	}

	// Test case 2: Overwrite an existing error code with a new ID
	existingCode := int32(http.StatusBadRequest)
	existingID := ids[existingCode]
	newID := "response.status.bad_request_custom" // Define a new error ID
	RegisterCode(existingCode, newID)
	if got, want := ids[existingCode], newID; got != want {
		t.Errorf("RegisterCode(%d, %s) overwrote existing code %d; got %s, want %s", existingCode, newID, existingCode, got, want)
	}
	// Restore the original ID for the sake of other tests
	ids[existingCode] = existingID
}

func TestError_Error(t *testing.T) {
	err := &Error{
		ID:     "response.status.bad_request",
		Code:   http.StatusBadRequest,
		Detail: "Bad request",
	}
	want := `{"id":"response.status.bad_request","code":400,"detail":"Bad request"}`
	if got := err.Error(); got != want {
		t.Errorf("Error() = %s; want %s", got, want)
	}
}

func TestError_String(t *testing.T) {
	err := &Error{
		ID:     "response.status.bad_request",
		Code:   http.StatusBadRequest,
		Detail: "Bad request",
	}
	want := `id:"response.status.bad_request" code:400 detail:"Bad request"`
	if got := err.String(); got != want {
		t.Errorf("String() = %s; want %s", got, want)
	}
}

func TestError_Status(t *testing.T) {
	err := &Error{
		ID:     "response.status.bad_request",
		Code:   http.StatusBadRequest,
		Detail: "Bad request",
	}
	want := http.StatusText(http.StatusBadRequest)
	if got := err.Status(); got != want {
		t.Errorf("Status() = %s; want %s", got, want)
	}
}

func TestError_Is(t *testing.T) {
	cases := []struct {
		err    error
		target *Error
		want   bool
	}{
		{
			err: &Error{
				ID:     "response.status.bad_request",
				Code:   http.StatusBadRequest,
				Detail: "Bad request",
			},
			target: &Error{
				Code: http.StatusBadRequest,
			},
			want: true,
		},
		{
			err: &Error{
				ID:     "response.status.bad_request",
				Code:   http.StatusBadRequest,
				Detail: "Bad request",
			},
			target: &Error{
				Code: http.StatusUnauthorized,
			},
			want: false,
		},
		{
			err:    errors.New("some error"),
			target: &Error{},
			want:   false,
		},
	}
	for _, c := range cases {
		var got *Error
		if errors.As(c.err, &got) && got.Is(c.target) != c.want {
			t.Errorf("Is(%v, %v) = %v; want %v", c.err, c.target, got, c.want)
		}
	}
}

func TestParse(t *testing.T) {
	cases := []struct {
		err  string
		want *Error
	}{
		{
			err: `{"id":"response.status.bad_request","code":400,"detail":"Bad request"}`,
			want: &Error{
				ID:     "response.status.bad_request",
				Code:   http.StatusBadRequest,
				Detail: "Bad request",
			},
		},
		{
			err: `{"id":"response.status.unauthorized","code":401,"detail":"Unauthorized"}`,
			want: &Error{
				ID:     "response.status.unauthorized",
				Code:   http.StatusUnauthorized,
				Detail: "Unauthorized",
			},
		},
		{
			err: "Some error",
			want: &Error{
				Detail: "Some error",
			},
		},
	}
	for _, c := range cases {
		got := Parse(c.err)
		if gotDetail := got.Detail; gotDetail != c.want.Detail {
			t.Errorf("Parse(%q).Detail = %q; want %q", c.err, gotDetail, c.want.Detail)
		}
		if gotID := got.ID; gotID != c.want.ID {
			t.Errorf("Parse(%q).ID = %q; want %q", c.err, gotID, c.want.ID)
		}
		if gotCode := got.Code; gotCode != c.want.Code {
			t.Errorf("Parse(%q).Code = %d; want %d", c.err, gotCode, c.want.Code)
		}
	}
}

func TestFromError(t *testing.T) {
	cases := []struct {
		err    error
		want   *Error
		wantOK bool
	}{
		{
			err: &Error{
				ID:     "response.status.bad_request",
				Code:   http.StatusBadRequest,
				Detail: "Bad request",
			},
			want: &Error{
				ID:     "response.status.bad_request",
				Code:   http.StatusBadRequest,
				Detail: "Bad request",
			},
			wantOK: true,
		},
		{
			err: errors.New("Some error"),
			want: &Error{
				Detail: "Some error",
			},
			wantOK: true,
		},
		{
			err:    NotFound("errors.rpc.test", "%s", "example"),
			want:   nil,
			wantOK: false,
		},
		{
			err: nil,
			want: &Error{
				Detail: "",
			},
			wantOK: true,
		},
	}
	for _, c := range cases {

		got := FromError(c.err)
		if got == nil && c.wantOK {
			t.Errorf("FromError(%v) = %v; want %v, %v", c.err, got, c.want, c.wantOK)
		}
		if gotDetail := got.Detail; gotDetail != c.want.Detail {
			t.Errorf("FromError(%v).Detail = %q; want %q", c.err, gotDetail, c.want.Detail)
		}
		if gotID := got.ID; gotID != c.want.ID {
			t.Errorf("FromError(%v).ID = %q; want %q", c.err, gotID, c.want.ID)
		}
		if gotCode := got.Code; gotCode != c.want.Code {
			t.Errorf("FromError(%v).Code = %d; want %d", c.err, gotCode, c.want.Code)
		}

	}
}

func TestBadRequest(t *testing.T) {
	type args struct {
		id     string
		format string
		obj    []any
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "Test with non-empty id and format",
			args: args{
				id:     "testID",
				format: "test format %s",
				obj:    []any{"test obj"},
			},
			want: &Error{
				ID:     "testID",
				Code:   http.StatusBadRequest,
				Detail: "test format test obj",
			},
		},
		{
			name: "Test with empty id and format",
			args: args{
				id:     "",
				format: "",
				obj:    []any{},
			},
			want: &Error{
				ID:     BadRequestID,
				Code:   http.StatusBadRequest,
				Detail: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BadRequest(tt.args.id, tt.args.format, tt.args.obj...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BadRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnauthorized(t *testing.T) {
	type args struct {
		id     string
		format string
		obj    []any
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "Test with non-empty id and format",
			args: args{
				id:     "testID",
				format: "test format %s",
				obj:    []any{"test obj"},
			},
			want: &Error{
				ID:     "testID",
				Code:   http.StatusUnauthorized,
				Detail: "test format test obj",
			},
		},
		{
			name: "Test with empty id and format",
			args: args{
				id:     "",
				format: "",
				obj:    []any{},
			},
			want: &Error{
				ID:     UnauthorizedID,
				Code:   http.StatusUnauthorized,
				Detail: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unauthorized(tt.args.id, tt.args.format, tt.args.obj...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unauthorized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForbidden(t *testing.T) {
	type args struct {
		id     string
		format string
		obj    []any
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "Test with non-empty id and format",
			args: args{
				id:     "testID",
				format: "test format %s",
				obj:    []any{"test obj"},
			},
			want: &Error{
				ID:     "testID",
				Code:   http.StatusForbidden,
				Detail: "test format test obj",
			},
		},
		{
			name: "Test with empty id and format",
			args: args{
				id:     "",
				format: "",
				obj:    []any{},
			},
			want: &Error{
				ID:     ForbiddenID,
				Code:   http.StatusForbidden,
				Detail: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Forbidden(tt.args.id, tt.args.format, tt.args.obj...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Forbidden() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotFound(t *testing.T) {
	err := NotFound("", "Resource not found")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	var notFoundErr *Error
	ok := errors.As(err, &notFoundErr)
	if !ok {
		t.Errorf("Expected type *Error, but got %T", err)
	}

	if notFoundErr.ID != NotFoundID {
		t.Errorf("Expected error ID %s, but got %s", NotFoundID, notFoundErr.ID)
	}

	if notFoundErr.Code != http.StatusNotFound {
		t.Errorf("Expected error code %d, but got %d", http.StatusNotFound, notFoundErr.Code)
	}

	if notFoundErr.Detail != "Resource not found" {
		t.Errorf("Expected error detail %s, but got %s", "Resource not found", notFoundErr.Detail)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	err := MethodNotAllowed("", "Method not allowed")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	var methodNotAllowedErr *Error
	ok := errors.As(err, &methodNotAllowedErr)
	if !ok {
		t.Errorf("Expected type *Error, but got %T", err)
	}

	if methodNotAllowedErr.ID != MethodNotAllowedID {
		t.Errorf("Expected error ID %s, but got %s", MethodNotAllowedID, methodNotAllowedErr.ID)
	}

	if methodNotAllowedErr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected error code %d, but got %d", http.StatusMethodNotAllowed, methodNotAllowedErr.Code)
	}

	if methodNotAllowedErr.Detail != "Method not allowed" {
		t.Errorf("Expected error detail %s, but got %s", "Method not allowed", methodNotAllowedErr.Detail)
	}
}

func TestTooManyRequests(t *testing.T) {
	err := TooManyRequests("", "Too many requests")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	var tooManyRequestsErr *Error
	ok := errors.As(err, &tooManyRequestsErr)
	if !ok {
		t.Errorf("Expected type *Error, but got %T", err)
	}

	if tooManyRequestsErr.ID != TooManyRequestsID {
		t.Errorf("Expected error ID %s, but got %s", TooManyRequestsID, tooManyRequestsErr.ID)
	}

	if tooManyRequestsErr.Code != http.StatusTooManyRequests {
		t.Errorf("Expected error code %d, but got %d", http.StatusTooManyRequests, tooManyRequestsErr.Code)
	}

	if tooManyRequestsErr.Detail != "Too many requests" {
		t.Errorf("Expected error detail %s, but got %s", "Too many requests", tooManyRequestsErr.Detail)
	}
}

func TestTimeout(t *testing.T) {
	err := Timeout("", "Request timeout")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	var timeoutErr *Error
	ok := errors.As(err, &timeoutErr)
	if !ok {
		t.Errorf("Expected type *Error, but got %T", err)
	}

	if timeoutErr.ID != RequestTimeoutID {
		t.Errorf("Expected error ID %s, but got %s", RequestTimeoutID, timeoutErr.ID)
	}

	if timeoutErr.Code != http.StatusRequestTimeout {
		t.Errorf("Expected error code %d, but got %d", http.StatusRequestTimeout, timeoutErr.Code)
	}

	if timeoutErr.Detail != "Request timeout" {
		t.Errorf("Expected error detail %s, but got %s", "Request timeout", timeoutErr.Detail)
	}
}

func TestConflict(t *testing.T) {
	err := Conflict("", "Resource conflict")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	var conflictErr *Error
	ok := errors.As(err, &conflictErr)
	if !ok {
		t.Errorf("Expected type *Error, but got %T", err)
	}

	if conflictErr.ID != ConflictID {
		t.Errorf("Expected error ID %s, but got %s", ConflictID, conflictErr.ID)
	}

	if conflictErr.Code != http.StatusConflict {
		t.Errorf("Expected error code %d, but got %d", http.StatusConflict, conflictErr.Code)
	}

	if conflictErr.Detail != "Resource conflict" {
		t.Errorf("Expected error detail %s, but got %s", "Resource conflict", conflictErr.Detail)
	}
}

func TestRequestEntityTooLarge(t *testing.T) {
	err := RequestEntityTooLarge("", "Request entity too large")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	var requestEntityTooLargeErr *Error
	ok := errors.As(err, &requestEntityTooLargeErr)
	if !ok {
		t.Errorf("Expected type *Error, but got %T", err)
	}

	if requestEntityTooLargeErr.ID != RequestEntityTooLargeID {
		t.Errorf("Expected error ID %s, but got %s", RequestEntityTooLargeID, requestEntityTooLargeErr.ID)
	}

	if requestEntityTooLargeErr.Code != http.StatusRequestEntityTooLarge {
		t.Errorf("Expected error code %d, but got %d", http.StatusRequestEntityTooLarge, requestEntityTooLargeErr.Code)
	}

	if requestEntityTooLargeErr.Detail != "Request entity too large" {
		t.Errorf("Expected error detail %s, but got %s", "Request entity too large", requestEntityTooLargeErr.Detail)
	}
}

func TestInternalServerError(t *testing.T) {
	err := InternalServerError("", "Internal server error")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	var internalServerErrorErr *Error
	ok := errors.As(err, &internalServerErrorErr)
	if !ok {
		t.Errorf("Expected type *Error, but got %T", err)
	}

	if internalServerErrorErr.ID != InternalServerErrorID {
		t.Errorf("Expected error ID %s, but got %s", InternalServerErrorID, internalServerErrorErr.ID)
	}

	if internalServerErrorErr.Code != http.StatusInternalServerError {
		t.Errorf("Expected error code %d, but got %d", http.StatusInternalServerError, internalServerErrorErr.Code)
	}

	if internalServerErrorErr.Detail != "Internal server error" {
		t.Errorf("Expected error detail %s, but got %s", "Internal server error", internalServerErrorErr.Detail)
	}
}

func TestNew(t *testing.T) {
	err := New("custom_id", 500, "Custom error")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	var newErr *Error
	ok := errors.As(err, &newErr)
	if !ok {
		t.Errorf("Expected type *Error, but got %T", err)
	}

	if newErr.ID != "custom_id" {
		t.Errorf("Expected error ID %s, but got %s", "custom_id", newErr.ID)
	}

	if newErr.Code != 500 {
		t.Errorf("Expected error code %d, but got %d", 500, newErr.Code)
	}

	if newErr.Detail != "Custom error" {
		t.Errorf("Expected error detail %s, but got %s", "Custom error", newErr.Detail)
	}
}

func TestNewf(t *testing.T) {
	err := NewFormat("custom_id", 500, "Custom error with %s", "data")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	var newfErr *Error
	ok := errors.As(err, &newfErr)
	if !ok {
		t.Errorf("Expected type *Error, but got %T", err)
	}

	if newfErr.ID != "custom_id" {
		t.Errorf("Expected error ID %s, but got %s", "custom_id", newfErr.ID)
	}

	if newfErr.Code != 500 {
		t.Errorf("Expected error code %d, but got %d", 500, newfErr.Code)
	}

	if newfErr.Detail != "Custom error with data" {
		t.Errorf("Expected error detail %s, but got %s", "Custom error with data", newfErr.Detail)
	}
}

func TestCode(t *testing.T) {
	err := Code(404, "Resource not found")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	var codeErr *Error
	ok := errors.As(err, &codeErr)
	if !ok {
		t.Errorf("Expected type *Error, but got %T", err)
	}

	if codeErr.ID != NotFoundID {
		t.Errorf("Expected error ID %s, but got %s", NotFoundID, codeErr.ID)
	}

	if codeErr.Code != 404 {
		t.Errorf("Expected error code %d, but got %d", 404, codeErr.Code)
	}

	if codeErr.Detail != "Resource not found" {
		t.Errorf("Expected error detail %s, but got %s", "Resource not found", codeErr.Detail)
	}
}
