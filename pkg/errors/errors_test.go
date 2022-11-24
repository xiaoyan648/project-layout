package errors

import (
	"fmt"
	"testing"
)

func TestWithCode(t *testing.T) {
	tests := []struct {
		code     int
		message  string
		wantType string
		wantCode int
	}{
		{ConfigurationNotValid, "ConfigurationNotValid error", "*WithCode", ConfigurationNotValid},
	}

	for _, tt := range tests {
		got := NewWithCode(tt.code, tt.message)
		if got.code != tt.wantCode {
			t.Errorf("WithCode(%v, %q): got: %v, want %v", tt.code, tt.message, got.code, tt.wantCode)
		}
	}
}

func TestWithCodef(t *testing.T) {
	tests := []struct {
		code       int
		format     string
		args       string
		wantType   string
		wantCode   int
		wangString string
	}{
		{ConfigurationNotValid, "Configuration %s", "failed", "*WithCode", ConfigurationNotValid, "ConfigurationNotValid error"},
	}

	for _, tt := range tests {
		got := NewWithCode(tt.code, tt.format, tt.args)
		if got.code != tt.wantCode {
			t.Errorf("WithCode(%v, %q %q): got: %v, want %v", tt.code, tt.format, tt.args, got.code, tt.wantCode)
		}

		if got.Error() != tt.wangString {
			t.Errorf("WithCode(%v, %q %q): got: %v, want %v", tt.code, tt.format, tt.args, got.Error(), tt.wangString)
		}
	}
}

func TestParseCoder(t *testing.T) {
	tests := []struct {
		name          string
		err           error
		wantHTTPCode  int
		wantString    string
		wantCode      int
		wantReference string
	}{
		{"std error", fmt.Errorf("std error"), 500, "An internal server error occurred", 1, "http://github.com/panda/errors/README.md"},
		{"WithCode error", NewWithCode(ErrInvalidJSON, "json error"), 500, "Data is not valid JSON", 1001, ""},
		{"grpc error", GRPCStatus(NewWithCode(ErrInvalidJSON, "json error")).Err(), 500, "Data is not valid JSON", 1001, ""},
		{"grpc error", GRPCStatus(fmt.Errorf("std error")).Err(), 500, "An internal server error occurred", 1, "http://github.com/panda/errors/README.md"},
	}

	for i, tt := range tests {
		coder := ParseCoder(tt.err)
		if coder.HTTPStatus() != tt.wantHTTPCode {
			t.Errorf("TestCodeParse(%d): got %q, want: %q", i, coder.HTTPStatus(), tt.wantHTTPCode)
		}

		if coder.String() != tt.wantString {
			t.Errorf("TestCodeParse(%d): got %q, want: %q", i, coder.String(), tt.wantString)
		}

		if coder.Code() != tt.wantCode {
			t.Errorf("TestCodeParse(%d): got %q, want: %q", i, coder.Code(), tt.wantCode)
		}

		if coder.Reference() != tt.wantReference {
			t.Errorf("TestCodeParse(%d): got %q, want: %q", i, coder.Reference(), tt.wantReference)
		}
	}

}
