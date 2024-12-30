// Code generated by "stringer -type=TokenType -trimprefix=ContextType -output=security_string.go"; DO NOT EDIT.

package security

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ContextTypeContext-0]
	_ = x[ContextTypeClientHeader-1]
	_ = x[ContextTypeServerHeader-2]
	_ = x[ContextTypeMetadataClient-3]
	_ = x[ContextTypeMetadata-4]
	_ = x[ContextTypeQuery-5]
	_ = x[ContextTypeCookie-6]
	_ = x[ContextTypeParam-7]
	_ = x[ContextTypeForm-8]
	_ = x[ContextTypeBody-9]
	_ = x[ContextTypeSession-10]
	_ = x[ContextTypeUnknown-11]
}

const _TokenType_name = "ContextClientHeaderServerHeaderMetadataClientMetadataQueryCookieParamFormBodySessionUnknown"

var _TokenType_index = [...]uint8{0, 7, 19, 31, 45, 53, 58, 64, 69, 73, 77, 84, 91}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
