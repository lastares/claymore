package jsonutil

import (
	"testing"
)

// TestJsonEncode tests the JsonEncode function for various input scenarios.
func TestJsonEncode(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:    "encode string",
			input:   "test string",
			want:    "\"test string\"",
			wantErr: false,
		},
		{
			name:    "encode integer",
			input:   123,
			want:    "123",
			wantErr: false,
		},
		{
			name:    "encode boolean",
			input:   true,
			want:    "true",
			wantErr: false,
		},
		{
			name:    "encode nil",
			input:   nil,
			want:    "null",
			wantErr: false,
		},
		{
			name:    "encode complex object",
			input:   struct{ Name string }{"Alice"},
			want:    "{\"Name\":\"Alice\"}",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonEncode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && string(got) != tt.want {
				t.Errorf("JsonEncode() got = %s, want %s", got, tt.want)
			}
		})
	}
}

// TestJsonDecode tests the JsonDecode function.
func TestJsonDecode(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    any
		wantErr bool
	}{
		{
			name:    "Valid JSON",
			input:   []byte(`{"key":"value"}`),
			want:    struct{ Key string }{Key: "value"},
			wantErr: false,
		},
		{
			name:    "Invalid JSON",
			input:   []byte(`{"key":"12345"}`),
			want:    struct{ Key string }{Key: "12345"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var v any = &struct {
				Key string
			}{}
			err := JsonDecode(tt.input, v)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && v != tt.want {
				t.Errorf("JsonDecode() got = %v, want %v", v, tt.want)
			}
		})
	}
}
